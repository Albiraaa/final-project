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

	//log.Println("Server running on :8080")
	//http.ListenAndServe(":8080", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, r)
}
