package main

import (
	"fmt"
	"net/http"
)

//routing register and login functionalities

func main() {
	http.HandleFunc("/auth/register", registerHandler)
	http.HandleFunc("/auth/login", loginHandler)

	fmt.Println("Authentication Service running on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
