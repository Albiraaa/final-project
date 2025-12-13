package main

import (
	"net/http"
	"os"
	"projek/config"
	"projek/routes"
)

func main() {
	config.ConnectDB()

	r := routes.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, r)
}
