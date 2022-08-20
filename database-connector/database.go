package databaseconnector

// This might help
// https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Construct DB entity, and return it, relatively simple. Can swap types of connectors more easily this way
// postgres://postgres:123456@127.0.0.1:5432/dummy
var postGresConnStr = os.Getenv("POSTGRES_DATABASE_URL")

var postGresDB, err = sql.Open("postgres", postGresConnStr)

// on localhost we need to disabled sslmode / I am lazy
// var postGresConnStr = "postgres://postgres:postgres@localhost:5432/timer?sslmode=disable"

// var mySQLConnStr = os.Getenv("MYSQL_DATABASE_URL")

func Init() {
	if postGresDB == nil {
		log.Fatal(err)
		panic("Database is down")
	}
}

func GetPostgresDatabaseHandler() *sql.DB {
	var postGresDB, err = sql.Open("postgres", postGresConnStr)
	if err != nil {
		log.Fatal(err)
		panic("Database is down")
	}

	return postGresDB
}

func CloseDBConnection(db *sql.DB) error {
	return db.Close()
}
