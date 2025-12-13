package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projek/config"
	"projek/models"

	"github.com/gorilla/mux"
)

func CreateLokasi(w http.ResponseWriter, r *http.Request) {
	var data models.Lokasi
	json.NewDecoder(r.Body).Decode(&data)
	config.DB.Create(&data)
	json.NewEncoder(w).Encode(data)
}

func GetLokasi(w http.ResponseWriter, r *http.Request) {
	var data []models.Lokasi
	config.DB.Find(&data)
	json.NewEncoder(w).Encode(data)
}

func GetLokasiByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var data models.Lokasi
	if config.DB.First(&data, id).RowsAffected == 0 {
		http.Error(w, "Lokasi tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func UpdateLokasi(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var data models.Lokasi
	if config.DB.First(&data, id).RowsAffected == 0 {
		http.Error(w, "Lokasi tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&data)
	config.DB.Save(&data)

	json.NewEncoder(w).Encode(data)
}

func DeleteLokasi(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if config.DB.Delete(&models.Lokasi{}, id).RowsAffected == 0 {
		http.Error(w, "Lokasi tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Lokasi berhasil dihapus",
	})
}
