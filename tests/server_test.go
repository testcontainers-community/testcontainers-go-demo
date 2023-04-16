package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"net/http"
	"net/http/httptest"
	"studentAPI/database"
	"studentAPI/server"
	"testing"
	"time"
)

func setupWithDockerFile(user, password, dbname string, ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context: "./custom",
			Dockerfile: "Dockerfile",
		},
		ExposedPorts: []string{"3306/tcp"}, //container port to expose
		Env: map[string]string{ 			//Values for container environmental variables
			"MARIADB_ROOT_PASSWORD": password,
			"MARIADB_USER":          user,
			"MARIADB_PASSWORD":      password,
			"MARIADB_DATABASE":      dbname,
		},
		//Checking Mariadb port for this string, to indicate container is fully up and running
		WaitingFor: wait.ForListeningPort("3306/tcp"),
	}

	//Starting the MariaDB Container
	customContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	//Stop tests if any errors encountered when setting up database connection
	if err != nil {
		return nil, err
	}

	return customContainer, nil
}

func TestB(t *testing.T) {
	user    := "theuser"
	password := "thepass"
	dbname   := "thedb"
	ctx := context.Background()
	log.Println(ctx)

	customContainer, err := setupWithDockerFile(user, password, dbname, ctx)

	// Get the container's host and port
	cHost, err := customContainer.Host(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	//obtaining the externally mapped port for the container
	cPort, err := customContainer.MappedPort(ctx, "3306")
	if err != nil {
		t.Fatal(err.Error())
	}

	//establish connection to MariaDB database container
	if err := database.Connect(user, password,cHost, cPort.Port(), dbname); err != nil {
		t.Fatal(err.Error())
	}

	//starting aAPI server to listen on port 9000
	router := server.SetupRouter()

	//Testing a POST HTTP request to /students
	t.Run(fmt.Sprintf("CREATE STUDENT API Test"), func(t *testing.T) {
		//Record to be submitted for insertion
		s := database.Student{
			Fname:       "Ben",
			Lname:       "Sterlin",
			DateOfBirth: time.Date(1998, time.August, 17, 23, 51, 42, 0, time.UTC),
			Email:       "Benlin@houses.org",
			Address:     "39 Benling Pass",
			Gender:      "Male",
		}

		//converting struct into json
		body, err := json.Marshal(s)
		if err != nil {
			t.Fatal("Unable to marshal student struct")
		}

		//Creating a GET HTTP request to /home
		req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		// initializing a recorder to keep track of response from the HTTP server
		w := httptest.NewRecorder()

		//Passing the request to the API router to serve
		router.ServeHTTP(w, req)

		//Using the assert sub package of the stretchr/testify to ensure the response was successful
		assert.Equal(t, http.StatusCreated, w.Code)

		// Asserting to ensure the response is body produced as expected
		assert.Equal(t, `{"success":true,"message":"","data":11}`, w.Body.String())
		t.Logf("Successfully created student record with response: %v \n", w.Body.String())
	})

	//Testing a GET HTTP request to /students/id
	t.Run(fmt.Sprintf("GET ONE BY ID Test"), func(t *testing.T) {
		// initializing a recorder to keep track of response from the HTTP server
		w := httptest.NewRecorder()

		//Creating a GET HTTP request to /home
		req, _ := http.NewRequest("GET", "/students/5", nil)

		//Passing the request to the API router to serve
		router.ServeHTTP(w, req)

		//Using the assert sub package of the stretchr/testify to ensure the response was successful
		assert.Equal(t, 200, w.Code)

		//Composing the expected response data
		expected:= server.Response{
			Success: true,
			Message: "",
			Data:    database.Student{
				ID:          5,
				Fname:       "Theda",
				Lname:       "Brockton",
				DateOfBirth: time.Date(1991, time.October, 29, 9, 8, 48, 0, time.UTC),
				Email:       "tbrockton4@lycos.com",
				Address:     "93 Hermina Plaza",
				Gender:      "Female",
			},
		}

		//converting expected data to json
		exprectedJson, err := json.Marshal(expected)
		if err != nil {
			t.Fatal(err.Error())
		}
		// Asserting to ensure the response is body produced as expected
		assert.Equal(t, string(exprectedJson), w.Body.String())
		t.Logf("Successfully retrieved with response: %s \n", w.Body.String())
	})
}


