package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/haifahrul/go-playground/internal/modules"
	"github.com/haifahrul/go-playground/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	/**
	 * Uncomment below if you want to use file .env
	 */
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	 * Uncomment below if you want to connection databse SQL
	 */
	// db.Conn, _ = helper.PostgreSql.Connect()

	fmt.Fprintf("Listing for requests at %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	err = http.ListenAndServe(
		":8080",
		router.Router,
	)

	if err != nil {
		log.Fatal("There was error when starting HTTP server", err)
	}
}
