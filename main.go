package main

import (
	"log"
	"net/http"

	"projek/config"
	"projek/routes"
)

func main() {
	config.ConnectDB()

	r := routes.InitRoutes()

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
