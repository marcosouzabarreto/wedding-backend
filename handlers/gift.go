package handlers

import (
	"log"
	"net/http"
	"wedding-backend/models"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
)

type GiftHandlers struct {
	service *services.GiftService
}

func NewGiftHandlers(service *services.GiftService) *GiftHandlers {
	return &GiftHandlers{service: service}
}

func (h *GiftHandlers) GetAll(c *gin.Context) {
	gifts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Gifts: %+v\n", gifts)
	c.JSON(http.StatusOK, gifts)
}

func (h *GiftHandlers) Create(c *gin.Context) {
	var gift models.Gift
	if err := c.ShouldBindJSON(&gift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	createdGift, err := h.service.Create(gift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdGift)
}

func (h *GiftHandlers) GetByID(c *gin.Context) {
	id := c.Param("id")
	gift, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "gift not found"})
		return
	}
	c.JSON(http.StatusOK, gift)
}

func (h *GiftHandlers) Update(c *gin.Context) {
	id := c.Param("id")
	var gift models.Gift
	if err := c.ShouldBindJSON(&gift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	updatedGift, err := h.service.Update(id, gift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedGift)
}

func (h *GiftHandlers) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gift deleted successfully"})
}
