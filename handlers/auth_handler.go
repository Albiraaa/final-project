package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"projek/config"
	"projek/models"

	"github.com/golang-jwt/jwt/v5"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	json.NewDecoder(r.Body).Decode(&input)
	config.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user)

	if user.ID == 0 {
		http.Error(w, "Login gagal", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("SECRET_KEY_HT"))
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
