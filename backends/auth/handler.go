// Implementation largely inspired by https://www.sohamkamani.com/golang/session-cookie-authentication/

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Credential struct models the structure of a user, both in the request body, and in the database schema
type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

// Each session contains a token, the username and the time at which it expires
type session struct {
	token    string
	username string
	expiry   time.Time
}

// isExpired determines if the session expiry field is before the current time and hence if the session has expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

// obtainCredentials by parsing and decoding the request body into a new `Credentials` instance
func obtainCredentials(requestBody io.ReadCloser) (Credentials, error) {
	creds := &Credentials{}
	err := json.NewDecoder(requestBody).Decode(creds)
	if err != nil {
		return *creds, errors.New("unable to decode credentials, invalid request body")
	}
	return *creds, nil
}

func Signup(w http.ResponseWriter, r *http.Request) {
	requestCredentials, err := obtainCredentials(r.Body)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Salt and hash the password using the bcrypt algorithm. The second argument is the cost of hashing, which we
	// arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(requestCredentials.Password), bcrypt.DefaultCost)

	// Insert the username, along with the hashed password into the database
	_, err = db.Query("insert into users values ($1, $2)", requestCredentials.Username, string(hashedPassword))
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent
	// back
}

func Signin(w http.ResponseWriter, r *http.Request) {
	requestCredentials, err := obtainCredentials(r.Body)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the existing entry present in the database for the given username
	result := db.QueryRow("select password from users where username=$1", requestCredentials.Username)

	// We create another instance of `Credentials` to store the credentials we get from the database
	storedCreds := &Credentials{}
	// Store the obtained password in `storedCreds`
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(requestCredentials.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}

	// Create a new random session token we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	if _, err = db.Query("insert into user_sessions values ($1, $2, $3)", sessionToken, requestCredentials.Username, expiresAt); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated we also set an
	// expiry time of 120 seconds
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
}

func IsAuth(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// Get session from database
	result := db.QueryRow("select username, token, expiry from user_sessions where token=$1", sessionToken) // need to explicitly state the order so that the correct values our matched with the struct

	storedSession := &session{}
	err = result.Scan(&storedSession.username, &storedSession.token, &storedSession.expiry)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If the session is present, but has expired, we can delete the session, and return an unauthorized status
	if storedSession.isExpired() {
		_, err := db.Exec("DELETE FROM user_sessions WHERE token=$1", sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If the session is valid, we return a 200
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// TODO: get session from db here
	result := db.QueryRow("select username, token, expiry from user_sessions where token=$1", sessionToken) // need to explicitly state the order so that the correct values our matched with the struct

	storedSession := &session{}
	err = result.Scan(&storedSession.username, &storedSession.token, &storedSession.expiry) // TODO: Need to find a way to scan the whole row into the struct so I can use the different values
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println(storedSession.expiry)
	// // We then get the session from our session map

	// // If the session is present, but has expired, we can delete the session, and return
	// // an unauthorized status
	if storedSession.isExpired() {
		_, err := db.Exec("DELETE FROM user_sessions WHERE token=$1", sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If the previous session is valid, create a new session token for the current user
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	if _, err = db.Query("UPDATE user_sessions SET token=$1, expiry=$2 WHERE username=$3", newSessionToken, expiresAt, storedSession.username); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

func Signout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// remove the users session from the session map
	_, err1 := db.Exec("DELETE FROM user_sessions WHERE token=$1", sessionToken)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time - this can be used in the UI to determine whethere or not to login (althout this is probably bad practice - a better way would be to have an isAuthenticated route that checks upon loading the website)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}
