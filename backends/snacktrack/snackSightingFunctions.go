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
type SnackSighting struct {
	SnackId          int    `json:"snackid" db:"snackid"`
	SighTime         string `json:"sighttime" db:"sighttime"`
	SightLocation    string `json:"sightlocation" db:"sightlocation"`
	SightEstDuration string `json:"sightestduration" db:"sightestduration"`
}

type SnackSightingKey struct {
	snackId       int    `json:"snackId"`
	sighTime      string `json:"sightTime"`
	sightLocation string `json:"sightLocation"`
}

func addSnackSighting(w http.ResponseWriter, r *http.Request) {
	requestSnackSighting, jsonError := obtainSnackSighting(r.Body)
	if jsonError != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.Query(` INSERT INTO snacksighting (snackid, sighttime, sightlocation, sightestduration)
							VALUES ($1, $2, $3, $4)`, requestSnackSighting.SnackId, requestSnackSighting.SighTime, requestSnackSighting.SightLocation, requestSnackSighting.SightEstDuration)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func readSnackSighting(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	result, err := db.Query("select snackid, sighttime, sightlocation, sightestduration from snacksighting")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		var snacksSightings []SnackSighting
		for result.Next() {
			sighting := &SnackSighting{}
			if err := result.Scan(&sighting.SnackId, &sighting.SighTime, &sighting.SightLocation, &sighting.SightEstDuration); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			//CHANGE TO ADD TO ARRAY AND RETURN ALL ARRAYS
			// var jsonOutput []byte
			fmt.Println(sighting.SightLocation)
			snacksSightings = append(snacksSightings, *sighting)
			// jsonOutput, err = json.Marshal(sighting)
			// fmt.Println(string(jsonOutput))
			// if err != nil {
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	return
			// }
			// w.Write(jsonOutput)
			// fmt.Fprintln(w, "Success!")
		}
		sightingsJson, _ := json.Marshal(snacksSightings)
		fmt.Println(string(sightingsJson))
		w.Write(sightingsJson)
	}
}
func removeSnackSighting(w http.ResponseWriter, r *http.Request) {
	requestSnackSightingKey, jsonError := obtainSnackSightingKey(r.Body)
	if jsonError != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.Query(` DELETE FROM snacksighting WHERE snackid=$1 AND sighttime=$2 AND sightlocation=$3)`,
		requestSnackSightingKey.snackId, requestSnackSightingKey.sighTime, requestSnackSightingKey.sightLocation)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func obtainSnackSighting(requestBody io.ReadCloser) (SnackSighting, error) {
	sighting := &SnackSighting{}
	err := json.NewDecoder(requestBody).Decode(sighting)
	if err != nil {
		return *sighting, errors.New("unable to decode snackSighting, invalid request body")
	}
	return *sighting, nil
}

func obtainSnackSightingKey(requestBody io.ReadCloser) (SnackSightingKey, error) {
	sightingKey := &SnackSightingKey{}
	err := json.NewDecoder(requestBody).Decode(sightingKey)
	if err != nil {
		return *sightingKey, errors.New("unable to decode snackSightingKey, invalid request body")
	}
	return *sightingKey, nil
}
