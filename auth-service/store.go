package main

import "sync"

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

var (
	users      = make(map[string]User)
	usersMutex sync.RWMutex //helps map in concurrent writing and access
	nextUserID = 1
)
