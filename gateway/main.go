package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	authURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		fmt.Println("Invalid Auth Service URL:", err)
		return
	}
	//Standard reverse proxy method provided by golang itself.
	authProxy := httputil.NewSingleHostReverseProxy(authURL)

	http.Handle("/auth/", authProxy)

	fmt.Println("API Gateway running on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gateway error:", err)
	}
}

