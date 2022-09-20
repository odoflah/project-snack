package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func addSnackSighting(w http.ResponseWriter, r *http.Request) {}
func readSnackSighting(w http.ResponseWriter, r *http.Request) {}
func removeSnackSighting(w http.ResponseWriter, r *http.Request) {}