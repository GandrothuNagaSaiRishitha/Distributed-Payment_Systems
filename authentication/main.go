package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authentication service is running")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request body, Please try registering again", http.StatusBadRequest)
		return
	}

	fmt.Println("Name:", req.Name)
	fmt.Println("Email:", req.Email)
	fmt.Println("Password:", req.Password)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"message": "Registration request received",
		"name":    req.Name,
		"email":   req.Email,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {

	conn := connectDB()

	if conn == nil {
		return
	}

	defer conn.Close(context.Background())

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/auth/register", registerHandler)

	fmt.Println("Authentication service running on :8080")

	http.ListenAndServe(":8080", nil)
}