package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"job-status-tracker/pkg/api"
	"job-status-tracker/pkg/database"
	"log"
	"net/http"
	"os"
)

func main() {
	// load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// read the datasource name from the environment variable
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// create a new database connection
	db, err := database.NewDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Create the router
	router := chi.NewRouter()

	// Set up middleware
	router.Use(middleware.Logger)

	// Define the routes
	router.Mount("/", api.Routes(db))

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
