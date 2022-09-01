package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	http.ListenAndServe(":8000", nil)
}

func getServiceHost(proxyConditionRaw string) (*url.URL, error) {
	var url *url.URL
	proxyCondition := strings.ToUpper(proxyConditionRaw)

	condition_url := os.Getenv(proxyCondition)

	if condition_url == "" {
		return url, errors.New("unknown service")
	}

	url, _ = url.Parse(condition_url)

	return url, nil
}

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	// ALGORITHM
	// - Get first elem in path -> determine the service
	// - do origin, _ := url.Parse("http://greeting:8002/") and pass the origin to the reverse proxy
	// remove the initial uri from the path and append the url.Path to be what remains

	path := r.URL.Path
	pathBreakdown := strings.Split(path, "/")
	serviceHost, _ := getServiceHost(pathBreakdown[1]) // the first uri in the path indifies the service

	// Mifght need to add this so it knows what kind of request it's dealing with cause it doesn;t know from teh incoming url
	// director := func(req *http.Request) {
	// 	req.Header.Add("X-Forwarded-Host", req.Host)
	// 	req.Header.Add("X-Origin-Host", origin.Host)
	// 	req.URL.Scheme = "http"
	// 	req.URL.Host = origin.Host
	// }

	// This is where I would be able to implement route protection for each host type individualy (for fine granulartity) or require full api authention (other than the autentication service which is an excpetion)

	// If we are autehticcate or trying to access the authAPI, then we can acess the backend services
	authService := "auth:8001"

	if serviceHost.Host == authService { // these requests will not have session cookies as they are requesting session tokens
		// acceptIngress
		reverseProxy := httputil.NewSingleHostReverseProxy(serviceHost)

		// reconstruct path
		var restOfUrl string
		for _, uri := range pathBreakdown[2:] {
			restOfUrl += uri + "/"
		}

		restOfUrlCorrected := strings.TrimSuffix(restOfUrl, "/")
		r.URL.Path = restOfUrlCorrected

		reverseProxy.ServeHTTP(w, r)
		return
	}

	// If not reqiuesting an authentication service - check auth cookies and see if user is authenticated
	cookieFromRequest, err := r.Cookie("session_token")

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

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+authService+"/isauth", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.AddCookie(cookieFromRequest)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(resp.Body)

	if resp.StatusCode == 200 {
		// if user is autheticated with appropriate session token cookies - connect them to requested service and serve
		reverseProxy := httputil.NewSingleHostReverseProxy(serviceHost)

		// reconstruct path
		var restOfUrl string
		for _, uri := range pathBreakdown[2:] {
			restOfUrl += uri + "/"
		}

		restOfUrlCorrected := strings.TrimSuffix(restOfUrl, "/")
		r.URL.Path = restOfUrlCorrected

		reverseProxy.ServeHTTP(w, r)
	} else {
		// unautheticated and not attempting to authenticate
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// func connectAndServe(serviceHost *url.URL, )
