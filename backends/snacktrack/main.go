package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

func main() {

	connectDb()
	migrateDb()

	http.HandleFunc("/submit-snack", submitSnack)
	http.HandleFunc("/get-snack", getSnack)
	http.HandleFunc("/remove-snack", removeSnack)
	http.HandleFunc("/submit-sighting", submitSighting)
	http.HandleFunc("/get-sightings", getSightings)
	http.HandleFunc("/remove-sighting", removeSighting)

	http.ListenAndServe(":80", nil)
}

func connectDb() {
	// Sleep to allow database process to start up...
	time.Sleep(5 * time.Second)

	// Create connection string based on environment (the production environment overwrites the development environment)
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"))

	// Open a connection to the database
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	// Verify connection to database
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func migrateDb() {
	// if db exists and schema matches current schema -> connect to db (need to connect to the bs server, not the individual database)
	// Create db
	// Create user
	// Upload schema
	// if db does not exist -> create database
	// if db exists but schema does not match updated schema -> update schema
	query, err := ioutil.ReadFile("./init.sql")
	if err != nil {
		panic(err)
	}
	sql := string(query)
	// fmt.Println("Printing sql")
	// fmt.Println(sql)
	if _, err := db.Exec(sql); err != nil {
		// panic(err)
		fmt.Println(err)
	}
}
