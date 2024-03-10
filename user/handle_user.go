package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
	if RepositoryErr != nil {
		http.Error(w, "Bad connection to SQL", http.StatusMethodNotAllowed)
	}

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
		ID:       uuid.New().String(),
		UserName: req.UserName,
	}
	userToSave := User{
		ID:       resp.ID,
		Username: resp.UserName,
	}

	err = UsersRepository.Save(userToSave)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	if RepositoryErr != nil {
		http.Error(w, "Bad connection to SQL", http.StatusMethodNotAllowed)
	}
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

	user, err := UsersRepository.FindByUsername(req.UserName)

	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
	}

	resp := loginUserResponse{
		URL: user.ID,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
