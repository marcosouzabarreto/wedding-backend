package handlers

import (
	"net/http"
	"wedding-backend/models"
	"wedding-backend/services"
	"github.com/gin-gonic/gin"
)

type RSVPHandler struct {
	service *services.RSVPService
}

func NewRSVPHandler(s *services.RSVPService) *RSVPHandler {
	return &RSVPHandler{service: s}
}

func (h *RSVPHandler) Create(c *gin.Context) {
	var rsvp models.RSVP
	if err := c.ShouldBindJSON(&rsvp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	createdRSVP, err := h.service.Create(&rsvp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRSVP)
}

func (h *RSVPHandler) CreateFamilyRSVP(c *gin.Context) {
	var req models.FamilyRSVPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	if err := h.service.CreateFamilyRSVP(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Family RSVP created successfully"})
}
