package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"wedding-backend/models"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
	"gorm.io/gorm"
)

type CreatePaymentRequest struct {
	GiftIDs      []uuid.UUID `json:"gift_ids"`
	CustomAmount float64     `json:"custom_amount"`
	GifterName   string      `json:"gifter_name"`
	Message      string      `json:"message"`
}

func CreatePayment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Req: %v", req)

		var gifts []models.Gift
		if len(req.GiftIDs) > 0 {
			if err := db.Where(req.GiftIDs).Find(&gifts).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch gifts"})
				return
			}
		}

		var total float64
		var giftNames []string
		for _, gift := range gifts {
			total += gift.Price
			giftNames = append(giftNames, gift.Name)
		}

		if req.CustomAmount > 0 {
			total += req.CustomAmount
		}

		if len(req.GiftIDs) > 0 || req.CustomAmount > 0 {
			userGiftService := services.NewUserGiftService(db)
			_, err := userGiftService.Create(req.GifterName, req.Message, req.GiftIDs, req.CustomAmount)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user gift"})
				return
			}
		}

		title := "Gifts for the couple"
		if len(giftNames) > 0 {
			title = fmt.Sprintf("Gifts for the couple: %s", strings.Join(giftNames, ", "))
		}

		if req.CustomAmount > 0 {
			if len(giftNames) > 0 {
				title = fmt.Sprintf("%s and a custom amount", title)
			} else {
				title = "Custom gift for the couple"
			}
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
					Title:       title,
					Description: "A collection of gifts for the happy couple",
					Quantity:    1,
					UnitPrice:   total,
				},
			},
			BackURLs: &preference.BackURLsRequest{
				Success: "https://laviekinho.com/gift-payment-status",
				Failure: "https://laviekinho.com/gift-payment-status",
				Pending: "https://laviekinho.com/gift-payment-status",
			},
		})

		log.Printf("Err: %v", err)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create preference"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"preferenceId": pref.ID})
	}
}
