package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

// SnackSighting struct models the structure of a SnackSighting entry, both in the request body, and in the database schema
type SnackSighting struct {
	SName         string `json:"sname" db:"sname"`
	SighTime      string `json:"sighttime" db:"sighttime"`
	SightLocation string `json:"sightlocation" db:"sightlocation"`
	SImage        string `json:"simage" db:"simage"`
}

// SnackSighting struct models the structure of a SnackSighting key, both in the request body, and in the database schema
type SnackSightingKey struct {
	SnackId       int    `json:"snackid" db:"snackid"`
	SighTime      string `json:"sighttime" db:"sighttime"`
	SightLocation string `json:"sightlocation" db:"sightlocation"`
}

func submitSighting(w http.ResponseWriter, r *http.Request) {
	requestSnackSighting, jsonError := obtainSnackSighting(r.Body)
	if jsonError != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.Query(`INSERT INTO snacksightings (sname, simage, sighttime, sightlocation) 
						VALUES ($1, $2, $3, $4)`,
		requestSnackSighting.SName, requestSnackSighting.SImage, requestSnackSighting.SighTime, requestSnackSighting.SightLocation)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func getSightings(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT sname, simage, sighttime, sightlocation FROM snacksightings")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		var snacksSightings []SnackSighting
		for result.Next() {
			//Append all returned entries to list
			sighting := &SnackSighting{}
			if err := result.Scan(&sighting.SName, &sighting.SImage, &sighting.SighTime, &sighting.SightLocation); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			snacksSightings = append(snacksSightings, *sighting)
		}
		//Convert list to JSON and write to http body
		sightingsJson, _ := json.Marshal(snacksSightings)
		w.Write(sightingsJson)
	}
}
func removeSighting(w http.ResponseWriter, r *http.Request) {
	requestSnackSightingKey, jsonError := obtainSnackSightingKey(r.Body)
	if jsonError != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.Query(`DELETE FROM snacksighting WHERE snackid=$1 AND sighttime=$2 AND sightlocation=$3`,
		requestSnackSightingKey.SnackId, requestSnackSightingKey.SighTime, requestSnackSightingKey.SightLocation)
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
