package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var db *sql.DB

var ctx = context.Background()

var server = "192.168.7.33\\SQLEXPRESS"
var user = "test"
var password = "test"
var database = "knowledgedb"

func init() {
	connect()
}

func connect() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;;",
		server, user, password, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	checkIfAlive()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
}

func checkIfAlive() error {
	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("No connectivity!")
		return err
	}
	return nil
}
