package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"studentAPI/database"
	"studentAPI/server"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "theuser"
	password = "thepass"
	dbname   = "thedb"
)

var db *sql.DB

func main() {
	err := database.Connect(user,password,host, port, dbname)
	if err != nil {
		log.Fatal(err)
	}
	//drops database table and records
	//database.Clear()

	//setup database table and insert few dummy records
	database.Setup()

	//start server to listen on port
	server.StartServer("9000")
}