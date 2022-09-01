package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

// Connect to database being used for authentication information
func dbConnect() {
	// Sleep to allow database process to start up...
	time.Sleep(5 * time.Second)

	// Create connection string based on environment (the production environment overwrites the development environment)
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"))

	// Open a connection to the database
	db, err := sql.Open("postgres", connectionString)
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

	// Initialise database connection
	fmt.Println("Initialising connection to database")
	dbConnect()
	defer db.Close()
	fmt.Println("Database connected successfully")

	// Start listing...
	http.ListenAndServe(":8001", nil)
}
