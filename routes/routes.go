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

	api.HandleFunc("/users/register", handlers.Register).Methods("POST")
	api.HandleFunc("/users/login", handlers.Login).Methods("POST")

	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/kategori", handlers.CreateKategori).Methods("POST")
	protected.HandleFunc("/kategori", handlers.GetKategori).Methods("GET")
	protected.HandleFunc("/kategori/{id}", handlers.GetKategoriByID).Methods("GET")
	protected.HandleFunc("/kategori/{id}", handlers.UpdateKategori).Methods("PUT")
	protected.HandleFunc("/kategori/{id}", handlers.DeleteKategori).Methods("DELETE")

	protected.HandleFunc("/lokasi", handlers.CreateLokasi).Methods("POST")
	protected.HandleFunc("/lokasi", handlers.GetLokasi).Methods("GET")
	protected.HandleFunc("/lokasi/{id}", handlers.GetLokasiByID).Methods("GET")
	protected.HandleFunc("/lokasi/{id}", handlers.UpdateLokasi).Methods("PUT")
	protected.HandleFunc("/lokasi/{id}", handlers.DeleteLokasi).Methods("DELETE")

	protected.HandleFunc("/ht", handlers.CreateHT).Methods("POST")
	protected.HandleFunc("/ht", handlers.GetHT).Methods("GET")
	protected.HandleFunc("/ht/{id}", handlers.GetHTByID).Methods("GET")
	protected.HandleFunc("/ht/{id}", handlers.UpdateHT).Methods("PUT")
	protected.HandleFunc("/ht/{id}", handlers.DeleteHT).Methods("DELETE")

	return r
}
