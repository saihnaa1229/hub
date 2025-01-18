package main

import (
	"log"
	"net/http"

	"hub/config"
	"hub/routes"

	_ "hub/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Hub API
// @version 1.0
// @description API documentation for the Hub project.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Connect to MongoDB
	config.ConnectDB()

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(r)

	// Add Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
