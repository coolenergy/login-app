package main

import (
	"github.com/Cerebrovinny/login-app/handlers"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Content-Type"},
	})

	handler := c.Handler(http.DefaultServeMux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
