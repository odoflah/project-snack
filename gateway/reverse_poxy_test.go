package main

import (
	"testing"
)

func TestServiceURLCreation(t *testing.T) {
	// TODO
	// Create user
	res, err := http.PostForm("http://localhost.com/auth/signup", url.Values{"username": {"Value"}, "id": {"123"}}) // make random string for username
	// Signin to get session token
	// Get greeting microservice once signed in
}

func TestInvalidServiceCall(t *testing.T) {
	// TODO
}
