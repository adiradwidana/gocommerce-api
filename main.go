package main

import (
	"github.com/gorilla/handlers"
	"hello/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.UserRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
