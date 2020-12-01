package accounts

import (
	"../core"
	"../db/models"
	"encoding/json"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	json.NewEncoder(w).Encode(user)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	err := AccountsService.registerUser(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("User has been registered successfully.")
}

func login(w http.ResponseWriter, r *http.Request) {
	user, token, err := AccountsService.login(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	data := core.Serialize(UserLoginSerializer{}, user).(UserLoginSerializer)
	data.Token = token
	json.NewEncoder(w).Encode(data)
}
