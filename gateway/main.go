package main

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// Create a reverse proxy connection to the backend service host, set the correct URL path for the specific service
// request and serve the request
func connectAndServe(serviceHost string, servicePath string, w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "OPTIONS" {
		return
	}
	// Convert the host to a URL and create a reverse proxy connection
	url, _ := url.Parse(serviceHost)
	reverseProxy := httputil.NewSingleHostReverseProxy(url)

	// Add the path for the service request
	r.URL.Path = servicePath

	// Serve the request
	reverseProxy.ServeHTTP(w, r)
}

// serviceHost based of the condition URI which is the first value in the path. If there is no service that corresponds
// to the URI then return ann error.
func serviceHost(proxyConditionRaw string) (string, error) {
	proxyCondition := strings.ToUpper(proxyConditionRaw)
	serviceHost := os.Getenv(proxyCondition)

	if serviceHost == "" {
		return "", errors.New("unknown service")
	}

	return serviceHost, nil
}

// constructServiceRequestURL by breaking down the path of the URL request to the gateway, extracting the host based of
// the first value in the path and then reconstructing the path from what remains.
func constructServiceRequestURL(requestPath string) (string, string) {
	// Break path down by forward slash separator
	pathBreakdown := strings.Split(requestPath, "/")

	// TODO: handle this error - maybe pass the error down to the handleRequest function
	// Obtain the appropriate host for the service based off the first value in the path
	serviceHost, _ := serviceHost(pathBreakdown[1])

	// Reconstruct the remainder of the path after having extracted the service URI
	var servicePath string
	for _, uri := range pathBreakdown[2:] {
		servicePath += uri + "/"
	}

	// Remove the trailing forward slash artifact from the path reconstruction
	servicePath = strings.TrimSuffix(servicePath, "/")

	return serviceHost, servicePath
}

// isAuthRequest creates a client to verify that the session cookie token is valid and the user has permissions to
// interface with the service. It returns the HTTP status code for the request, a 200
func isAuthRequest(cookieFromRequest *http.Cookie) (bool, error) {
	// Create client for request
	client := &http.Client{}

	// Construct request to authentication service
	req, err := http.NewRequest("GET", os.Getenv("AUTH")+"/isauth", nil)
	if err != nil {
		return false, errors.New("unable to construct request to authentication service")
	}

	// Add cookie from request to proxy server to the new request to the authentication service and make request
	req.AddCookie(cookieFromRequest)
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("unable to make request to authentication service")
	}

	// If the authentication service returns success status code the session token in the cookie is valid, else return
	// an error
	if resp.StatusCode == 200 {
		return true, nil
	} else {
		return false, errors.New("bad status in authentication service response")
	}
}

// handleRequest by constructing the URL for the service request. If the incoming request is a request to the
// authentication service, then connect to the service and serve the request. Otherwise, if it is a request to another
// service, check if authenticated by a valid session, connect to the service and serve the request. If no session
// cookie or an invalid session is provided, return an unauthorized response. Route protection is implemented at this
// level.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "OPTIONS" {
		return
	}
	requestPath := r.URL.Path
	serviceHost, servicePath := constructServiceRequestURL(requestPath)

	// If we trying to access the authentication service, then let us. These requests will not have session cookies as
	// they are requesting session tokens.
	if serviceHost == os.Getenv("AUTH") {
		connectAndServe(serviceHost, servicePath, w, r)
		return // BUG: this is not needed?
	}

	// If not requesting an authentication service, check cookie to see see if user is authenticated
	cookieFromRequest, err := r.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest) // unable to process the request
		return
	}

	// Make a request to check if user session token in cookie is valid
	isAuthRequestStatus, err := isAuthRequest(cookieFromRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isAuthRequestStatus {
		// If user is authenticated with appropriate session token cookie, connect them to requested service and serve the request
		connectAndServe(serviceHost, servicePath, w, r)
		return // BUG: this is not needed?
	} else {
		// If unauthenticated AND not attempting to authenticate (checked at the beginning of the function), then return
		// an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func enableCors(w http.ResponseWriter) {
	//Allow CORS here By * or specific origin
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// API Gateway (reverse-proxy) entrypoint
func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8000", nil)
}
