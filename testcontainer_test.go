package main

import (
	"context"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
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
	var err error
	mariadbContainer, err = mariadb.RunContainer(ctx,
		testcontainers.WithImage("mariadb:10.6"),
		mariadb.WithDatabase(dbname),
		mariadb.WithUsername(user),
		mariadb.WithPassword(password),
	)

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
