package main

import (
	"errors"
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
	// origin, _ := url.Parse("http://greeting:8002/")
	// ALGORITHM
	// - Get first elem in path -> determine the service
	// - do origin, _ := url.Parse("http://greeting:8002/") and pass the origin to the reverse proxy
	// remove the initial uri from the path and append the url.Path to be what remains

	path := r.URL.Path
	pathBreakdown := strings.Split(path, "/")
	serviceHost, _ := getServiceHost((pathBreakdown[1])) // the first uri in the path indifies the service

	// Mifght need to add this so it knows what kind of request it's dealing with cause it doesn;t know from teh incoming url
	// director := func(req *http.Request) {
	// 	req.Header.Add("X-Forwarded-Host", req.Host)
	// 	req.Header.Add("X-Origin-Host", origin.Host)
	// 	req.URL.Scheme = "http"
	// 	req.URL.Host = origin.Host
	// }

	reverseProxy := httputil.NewSingleHostReverseProxy(serviceHost)

	// reconstruct path
	var restOfUrl string
	for _, uri := range pathBreakdown[2:] {
		restOfUrl += uri + "/"
	}

	restOfUrlCorrected := strings.TrimSuffix(restOfUrl, "/")
	r.URL.Path = restOfUrlCorrected

	reverseProxy.ServeHTTP(w, r)
}
