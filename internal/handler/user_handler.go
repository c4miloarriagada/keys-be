package handler

import (
	"net/http"
	"strconv"

	"github.com/c4miloarriagada/keys-be/internal/domain"
	domain_errors "github.com/c4miloarriagada/keys-be/internal/domain/errors"
	"github.com/c4miloarriagada/keys-be/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		writeError(c, domain_errors.NewValidationError("invalid_id", "invalid id"))
		return
	}

	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		writeError(c, err)
		return
	}

	if user == nil {
		writeError(c, domain_errors.NewNotFoundError("user_not_found", "user not found"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		writeError(c, domain_errors.NewValidationError("invalid_body", "invalid body"))
		return
	}

	if err := h.UserService.CreateUser(&user); err != nil {
		writeError(c, domain_errors.NewInternalServerError("internal_error", "error creating user"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado correctamente"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		writeError(c, domain_errors.NewInternalServerError("internal_error", "error getting users"))
		return
	}

	c.JSON(http.StatusOK, users)
}
