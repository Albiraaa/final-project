package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projek/config"
	"projek/models"

	"github.com/gorilla/mux"
)

func CreateHT(w http.ResponseWriter, r *http.Request) {
	var data models.HT
	json.NewDecoder(r.Body).Decode(&data)
	config.DB.Create(&data)
	json.NewEncoder(w).Encode(data)
}

func GetHT(w http.ResponseWriter, r *http.Request) {
	var data []models.HT
	config.DB.Preload("Kategori").Preload("Lokasi").Find(&data)
	json.NewEncoder(w).Encode(data)
}

func GetHTByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var data models.HT
	result := config.DB.Preload("Kategori").Preload("Lokasi").First(&data, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Data HT tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func UpdateHT(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var data models.HT
	if config.DB.First(&data, id).RowsAffected == 0 {
		http.Error(w, "Data HT tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&data)
	config.DB.Save(&data)

	json.NewEncoder(w).Encode(data)
}

func DeleteHT(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	result := config.DB.Delete(&models.HT{}, id)
	if result.RowsAffected == 0 {
		http.Error(w, "Data HT tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data HT berhasil dihapus",
	})
}
