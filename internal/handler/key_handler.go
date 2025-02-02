package handler

import (
	"log"

	domain_errors "github.com/c4miloarriagada/keys-be/internal/domain/errors"
	"github.com/c4miloarriagada/keys-be/internal/service"
	"github.com/gin-gonic/gin"
)

type KeyHandler struct {
	KeyService *service.KeyService
}

func NewKeyHandler(service *service.KeyService) *KeyHandler {
	return &KeyHandler{KeyService: service}
}

func (h *KeyHandler) Save(c *gin.Context) {
	var keyDTO keyDTO
	if err := c.BindJSON(&keyDTO); err != nil {
		log.Print("Error binding json: ", err)
		writeError(c, domain_errors.NewInternalServerError("internal_error", "error binding json"))
		return
	}

	key := keyDTO.toDomain()
	if err := h.KeyService.Save(&key); err != nil {
		log.Print("Error saving key: ", err)
		writeError(c, err)
		return
	}

	c.JSON(201, key)

}
