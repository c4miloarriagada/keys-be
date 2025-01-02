package http

import (
	"encoding/json"
	"net/http"

	"github.com/c4miloarriagada/keys-be/internal/core/service"
)

type APIHandler struct {
	UserService *service.UserService
}

func NewAPIHandler(us *service.UserService) *APIHandler {
	return &APIHandler{UserService: us}
}

func (h *APIHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
