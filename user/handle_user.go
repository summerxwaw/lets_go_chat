package user

import (
	"encoding/json"
	"net/http"
)

type createUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type createUserResponse struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
}

type loginUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	URL string `json:"url"`
}

func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method ", http.StatusMethodNotAllowed)
		return
	}

	var req createUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	resp := createUserResponse{
		ID:       "some-unique-uuid",
		UserName: req.UserName,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req loginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	resp := loginUserResponse{
		URL: "USER - TOKEN",
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
