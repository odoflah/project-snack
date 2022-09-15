package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

func configureDatabase() {
	// if db exists and schema matches current schema -> connect to db (need to connect to the bs server, not the individual database)
		// Create db
		// Create user
		// Upload schema
	// if db does not exist -> create database
	// if db exists but schema does not match updated schema -> update schema
}

// Connect to database being used for authentication information
func dbConnect() {
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

func main() {
	// Authentication service routes
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/signout", Signout)
	http.HandleFunc("/isauth", IsAuth)
	// TODO: Add return auth user data e.g. returns names, userid...

	// Initialise database connection
	fmt.Println("Initialising connection to database")
	dbConnect()
	defer db.Close()

	fmt.Println("Database connected successfully")

	// Start listing...
	log.Fatalln(http.ListenAndServe(":80", nil))
}
