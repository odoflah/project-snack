package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Get the url for a given proxy condition
func getProxyUrl(proxyConditionRaw string) string {
	proxyCondition := strings.ToUpper(proxyConditionRaw)

	condition_url := os.Getenv(proxyCondition)

	if condition_url == "" {
		return "http://gateway:8000/invalid"
	}

	return condition_url
}

/*
	Reverse Proxy Logic
*/

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	// proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	fmt.Println("I'm here")

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	// proxy.ServeHTTP(res, req)
	// TODO: make an http request here to test if I can make a request
	_, err := http.Get("http://greeting:8002")
	if err != nil {
		log.Fatalln(err)
	}
}

func service(url string) string {
	split := strings.Split(url, string(os.PathSeparator))
	return split[1]
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	// requestPayload := parseRequestBody(req)
	path := req.URL.Path
	// fmt.Println(path)
	service := service(path)
	// fmt.Println(service)

	url := getProxyUrl(service)

	fmt.Println(url)

	serveReverseProxy(url, res, req)
}

/*
	Entry
*/

func invalid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	// Log setup values
	// logSetup()

	// start server
	http.HandleFunc("/invalid", invalid)
	http.HandleFunc("/", handleRequestAndRedirect)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
