package handlers

import (
	"net/http"
	"wedding-backend/models"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
)

type GuestHandlers struct {
	service *services.GuestService
}

func NewGuestHandlers(service *services.GuestService) *GuestHandlers {
	return &GuestHandlers{service: service}
}

func (h *GuestHandlers) GetAll(c *gin.Context) {
	guests, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guests)
}

func (h *GuestHandlers) Create(c *gin.Context) {
	var guest models.Guest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	createdGuest, err := h.service.Create(guest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdGuest)
}

func (h *GuestHandlers) GetByID(c *gin.Context) {
	id := c.Param("id")
	guest, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "guest not found"})
		return
	}
	c.JSON(http.StatusOK, guest)
}

func (h *GuestHandlers) Update(c *gin.Context) {
	id := c.Param("id")
	var guest models.Guest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	updatedGuest, err := h.service.Update(id, &guest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedGuest)
}