package main

import (
	"os"
	"log"
	"encoding/json"
	"net/url"
	"net/http"
	"net/http/httputil"
)

var (
	params = os.Getenv("PARAMS")
	hostname = os.Getenv("HOSTNAME")
	port = os.Getenv("PORT")
)

func validateConfiguration() {
	if hostname == "" {
		log.Fatal("Hostname is required")
	}
	if port  == "" {
		log.Fatal("Port number is missing")
	}
	_, err := url.ParseRequestURI(hostname) // doesn't provide an absolute validation.
	if err != nil {
		log.Fatal("Hostname is invalid")
	}
	log.Printf("Redirecting to %s\n", hostname)
}

func extendQueryParams(req *http.Request) {
	var queryParams map[string]string
	err := json.Unmarshal([]byte(params), &queryParams)
	if err != nil {
		log.Fatal("Invalid parameters format")
	}
	q := req.URL.Query()
	for k, v := range queryParams {
		if q.Get(k) == "" {
			q.Add(k, v)
		}
		// in case if we want to overwrite parameters, we can uncomment code below and comment out code above 
		//q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received")
	url, err := url.Parse(hostname)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Host = url.Host

	if params != "" {
		extendQueryParams(req)
		log.Printf("Query parameters %s\n", req.URL.RawQuery)
	}

	proxy.ServeHTTP(res, req)
}

func runProxy() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func main() {
	validateConfiguration()
	runProxy()
}
