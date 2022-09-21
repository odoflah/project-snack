package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	// "database/sql"
	// "os"
	// "time"

	_ "github.com/lib/pq"
)

// Credential struct models the structure of a user, both in the request body, and in the database schema
type Snack struct {
	SnackId int `json:"snackid" db:"snackid"`
	SnackName string `json:"snackname" db:"snackname"`
	SnackDesc string `json:"snackdesc" db:"snackdes"`
	SnackCat string `json:"snackcat" db:"snackcat"`
	SnackPic string `json:"snackpic" db:"snackpic"`
	HealthScore int `json:"healthscore" db:"healthscore"`
}

type SnackKey struct {
	SnackId          int    `json:"snackid" db:"snackid"`
}

func submitSnack(w http.ResponseWriter, r *http.Request) {
	requestSnack, jsonError := obtainSnack(r.Body)
	fmt.Println(requestSnack.SnackName)
	if jsonError != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.Query(` INSERT INTO snacks (snackname, snackdesc, snackcat, snackpic, healthscore)
							VALUES ($1, $2, $3, $4, $5)`, requestSnack.SnackName, requestSnack.SnackDesc, requestSnack.SnackCat, requestSnack.SnackPic, requestSnack.HealthScore)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func getSnack(w http.ResponseWriter, r *http.Request) {
	requestSnackKey, jsonError := obtainSnackKey(r.Body)
	if jsonError != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := db.Query("select snackid, snackname, snackdesc, snackcat, snackpic, healthscore from snacks where snackid=$1", requestSnackKey.SnackId )
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		var snacks []Snack
		for result.Next() {
			snack := &Snack{}
			if err := result.Scan(&snack.SnackId, &snack.SnackName, &snack.SnackDesc, &snack.SnackCat, &snack.SnackPic, &snack.HealthScore); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			snacks = append(snacks, *snack)
		}
		snacksJson, _ := json.Marshal(snacks)
		fmt.Println(string(snacksJson))
		w.Write(snacksJson)
	}
}
func removeSnack(w http.ResponseWriter, r *http.Request) {
	requestSnackKey, jsonError := obtainSnackKey(r.Body)
	if jsonError != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.Query(` DELETE FROM snacks WHERE snackid=$1`, requestSnackKey.SnackId)
	if err != nil {
		// If there is any issue with removing from the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func obtainSnack(requestBody io.ReadCloser) (Snack, error) {
	snack := &Snack{}
	err := json.NewDecoder(requestBody).Decode(snack)
	if err != nil {
		fmt.Println(err)
		return *snack, errors.New("unable to decode Snack, invalid request body")
	}
	return *snack, nil
}

func obtainSnackKey(requestBody io.ReadCloser) (SnackKey, error) {
	snackKey := &SnackKey{}
	err := json.NewDecoder(requestBody).Decode(snackKey)
	if err != nil {
		fmt.Println(err)
		return *snackKey, errors.New("unable to decode SnackKey, invalid request body")
	}
	return *snackKey, nil
}
