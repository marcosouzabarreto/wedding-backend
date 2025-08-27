package handlers

import (
	"net/http"
	"wedding-backend/models"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
)

type FamilyHandlers struct {
	service *services.FamilyService
}

func NewFamilyHandlers(service *services.FamilyService) *FamilyHandlers {
	return &FamilyHandlers{service: service}
}

func (h *FamilyHandlers) GetAll(c *gin.Context) {
	families, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, families)
}

func (h *FamilyHandlers) Create(c *gin.Context) {
	var familyInput models.FamilyInput
	if err := c.ShouldBindJSON(&familyInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	family, err := h.service.Create(familyInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, family)
}

func (h *FamilyHandlers) GetById(c *gin.Context) {
	id := c.Param("id")
	family, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "family not found"})
		return
	}
	c.JSON(http.StatusOK, family)
}

func (h *FamilyHandlers) GetByToken(c *gin.Context) {
	token := c.Param("token")
	family, err := h.service.GetByToken(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "family not found"})
		return
	}
	c.JSON(http.StatusOK, family)
}
