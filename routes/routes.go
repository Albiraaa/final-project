package routes

import (
	"github.com/gorilla/mux"

	"projek/handlers"
	"projek/middleware"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CORSMiddleware)

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/users/login", handlers.Login).
		Methods("POST", "OPTIONS")

	api.HandleFunc("/users/register", handlers.Register).
		Methods("POST", "OPTIONS")

	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/kategori", handlers.CreateKategori).Methods("POST", "OPTIONS")
	protected.HandleFunc("/kategori", handlers.GetKategori).Methods("GET", "OPTIONS")
	protected.HandleFunc("/kategori/{id}", handlers.GetKategoriByID).Methods("GET", "OPTIONS")
	protected.HandleFunc("/kategori/{id}", handlers.UpdateKategori).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/kategori/{id}", handlers.DeleteKategori).Methods("DELETE", "OPTIONS")

	protected.HandleFunc("/lokasi", handlers.CreateLokasi).Methods("POST", "OPTIONS")
	protected.HandleFunc("/lokasi", handlers.GetLokasi).Methods("GET", "OPTIONS")
	protected.HandleFunc("/lokasi/{id}", handlers.GetLokasiByID).Methods("GET", "OPTIONS")
	protected.HandleFunc("/lokasi/{id}", handlers.UpdateLokasi).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/lokasi/{id}", handlers.DeleteLokasi).Methods("DELETE", "OPTIONS")

	protected.HandleFunc("/ht", handlers.CreateHT).Methods("POST", "OPTIONS")
	protected.HandleFunc("/ht", handlers.GetHT).Methods("GET", "OPTIONS")
	protected.HandleFunc("/ht/{id}", handlers.GetHTByID).Methods("GET", "OPTIONS")
	protected.HandleFunc("/ht/{id}", handlers.UpdateHT).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/ht/{id}", handlers.DeleteHT).Methods("DELETE", "OPTIONS")

	return r
}
