package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//encode user slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Error encoding users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var payloadUser User
	err := json.NewDecoder(r.Body).Decode(&payloadUser)
	users = append(users, User{Id: payloadUser.Id, Name: payloadUser.Name})

	//encode user slice to json

	if err != nil {
		http.Error(w, "Error encoding users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
