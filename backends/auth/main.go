package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

func main() {
	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/signout", Signout)
	http.HandleFunc("/isauth", IsAuth)
	// initialize our database connection
	fmt.Println("Initialising connection to database")
	dbConnect()
	defer db.Close()
	fmt.Println("Database connected sucessfullly")
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func dbConnect() {
	var err error
	time.Sleep(5 * time.Second)
	connStr := "user=dev dbname=dev_db password=password host=db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}
