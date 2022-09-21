package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"os"
	"time"
	"io/ioutil"

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

	migrateDb()

	http.HandleFunc("/", hello)

	http.HandleFunc("/addSnack", addSnack)
	http.HandleFunc("/readSnack", readSnack)
	http.HandleFunc("/removeSnack", removeSnack)
	http.HandleFunc("/addSnackSighting", addSnackSighting)
	http.HandleFunc("/readSnackSighting", readSnackSighting)
	http.HandleFunc("/removeSnackSighting", removeSnackSighting)

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