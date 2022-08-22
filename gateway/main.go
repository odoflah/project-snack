package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// Get the url for a given proxy condition
func getProxyUrl(proxyConditionRaw string) string {
	proxyCondition := strings.ToUpper(proxyConditionRaw)

	condition_url := os.Getenv(proxyCondition)

	if condition_url == "" {
		return "invalid"
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
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	fmt.Println("I'm here")

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
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

	if url == "invalid" {
		res.WriteHeader(http.StatusBadRequest)
	} else {
		serveReverseProxy(url, res, req)
	}
}

/*
	Entry
*/

func main() {
	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
