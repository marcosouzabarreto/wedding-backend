package handlers

import (
	"net/http"
	"wedding-backend/models"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
)

type UserGiftHandlers struct {
	service *services.UserGiftService
}

func NewUserGiftHandlers(service *services.UserGiftService) *UserGiftHandlers {
	return &UserGiftHandlers{service: service}
}

func (h *UserGiftHandlers) GetAll(c *gin.Context) {
	userGifts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userGifts)
}

func (h *UserGiftHandlers) Create(c *gin.Context) {
	var userGift models.UserGift
	if err := c.ShouldBindJSON(&userGift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	createdUserGift, err := h.service.Create(userGift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUserGift)
}
