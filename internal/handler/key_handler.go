package handler

import (
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	key := keyDTO.toDomain()
	if err := h.KeyService.Save(&key); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, key)

}
