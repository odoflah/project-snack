package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)
// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

func main() {
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

	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is snacktrack!")
	db.Query(` INSERT INTO snacks (snackname, snackdesc, snackcat, snackpic, healthscore)
							VALUES ('test', 'test', 'test', 'test', 5)`)
	result, err := db.Query("select * from snacks")
	if err != nil{
		fmt.Fprintln(w, err)
	}else{
		for result.Next() {
			var id int
			var name string
			var desc string
			var cat string
			var image string
			var score int

			if err := result.Scan(&id, &name, &desc, &cat, &image, &score); err != nil {
				fmt.Fprintln(w, err)
			}
			fmt.Fprintln(w, name)
		}
	}	
}
