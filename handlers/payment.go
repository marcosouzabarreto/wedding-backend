package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"wedding-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
	"gorm.io/gorm"
)

type CreatePaymentRequest struct {
	GiftIDs []uint `json:"gift_ids"`
}

func CreatePayment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var gifts []models.Gift
		if err := db.Where(req.GiftIDs).Find(&gifts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch gifts"})
			return
		}

		var total float64
		var giftNames []string
		for _, gift := range gifts {
			total += gift.Price
			giftNames = append(giftNames, gift.Name)
		}

		cfg, err := config.New(os.Getenv("MERCADOPAGO_ACCESS_TOKEN"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to configure mercadopago"})
			return
		}

		client := preference.NewClient(cfg)

		pref, err := client.Create(context.Background(), preference.Request{
			Items: []preference.ItemRequest{
				{
					Title:       fmt.Sprintf("Gifts for the couple: %s", strings.Join(giftNames, ", ")),
					Description: "A collection of gifts for the happy couple",
					Quantity:    1,
					UnitPrice:   total,
				},
			},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create preference"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"preferenceId": pref.ID})
	}
}
