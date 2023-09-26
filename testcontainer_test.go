package main

import (
	"context"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"studentAPI/database"
	"testing"
)

// Declaration of global variables
var (
	mariadbContainer testcontainers.Container
	mariadbHost      string
	mariadbPort      nat.Port
	ctx              = context.Background()
)

func setupContainer() error {
	// Defining parameters for MariaDB container
	req := testcontainers.ContainerRequest{
		Image:        "mariadb:10.6",       //mariadb version version
		ExposedPorts: []string{"3306/tcp"}, //container port to expose
		Env: map[string]string{ //Values for container environmental variables
			"MARIADB_ROOT_PASSWORD": password,
			"MARIADB_USER":          user,
			"MARIADB_PASSWORD":      password,
			"MARIADB_DATABASE":      dbname,
		},
		SkipReaper: true,
		//Checking Mariadb port for this string, to indicate container is fully up and running
		WaitingFor: wait.ForListeningPort("3306/tcp"),
	}

	//Starting the MariaDB Container
	var err error
	mariadbContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	//Stop tests if any errors encountered when setting up database connection
	if err != nil {
		return err
	}

	// Get the container's host and port
	mariadbHost, err = mariadbContainer.Host(ctx)
	if err != nil {
		return err
	}
	//obtaining the externally mapped port for the container
	mariadbPort, err = mariadbContainer.MappedPort(ctx, "3306")
	if err != nil {
		return err
	}
	return nil
}

// Perform database migration
func setupDBConnection() error {
	//establish connection to MariaDB database container
	err := database.Connect(user, password, mariadbHost, mariadbPort.Port(), dbname)
	if err != nil {
		return err
	}

	//setup student table and insert few dummy records
	database.Setup()
	return nil
}

// This contains setup and teardown code gets called before all test functions.
func TestMain(m *testing.M) {
	//Set up container

	log.Println("herererere")

	if err := setupContainer(); err != nil {
		log.Fatal(err.Error())
	}

	//Run DB Migrations
	if err := setupDBConnection(); err != nil {
		log.Fatal(err.Error())
	}

	//executing all other test suite
	exitCode := m.Run()

	//Destruct database container after completing tests
	if err := mariadbContainer.Terminate(ctx); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
	os.Exit(exitCode)
}
