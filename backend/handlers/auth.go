package handlers

import (
	"encoding/json"
	"net/http"

	"backend/utils"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Autenticação simplificada para demonstração
	if creds.Username == "user" && creds.Password == "password" {
		token, err := utils.GenerateJWT(creds.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"token": token})
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
