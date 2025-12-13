package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projek/config"
	"projek/models"

	"github.com/gorilla/mux"
)

func CreateKategori(w http.ResponseWriter, r *http.Request) {
	var data models.Kategori
	json.NewDecoder(r.Body).Decode(&data)
	config.DB.Create(&data)
	json.NewEncoder(w).Encode(data)
}

func GetKategori(w http.ResponseWriter, r *http.Request) {
	var data []models.Kategori
	config.DB.Find(&data)
	json.NewEncoder(w).Encode(data)
}

func GetKategoriByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var data models.Kategori
	if config.DB.First(&data, id).RowsAffected == 0 {
		http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func UpdateKategori(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var data models.Kategori
	if config.DB.First(&data, id).RowsAffected == 0 {
		http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&data)
	config.DB.Save(&data)

	json.NewEncoder(w).Encode(data)
}

func DeleteKategori(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if config.DB.Delete(&models.Kategori{}, id).RowsAffected == 0 {
		http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Kategori berhasil dihapus",
	})
}
