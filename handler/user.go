package handler

import (
	"encoding/json"
	"net/http"
)

// UserHandler struct
type UserHandler struct {
}

// NewUserHandler constructor
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type UserResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (handle *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := UserResponse{
		Status:  "ok",
		Version: "1.0.0",
	}
	w.Header().Set("GET /listuser", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
