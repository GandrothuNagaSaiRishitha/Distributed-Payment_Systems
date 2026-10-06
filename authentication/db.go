package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func connectDB() *pgx.Conn {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Could not load .env")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		fmt.Println("Could not connect to database:", err)
		return nil
	}

	fmt.Println("Connected to PostgreSQL!")

	return conn
}

func createUser(conn *pgx.Conn, name string, email string, passwordHash string) error {

	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)",
		name,
		email,
		passwordHash,
	)

	return err
}

