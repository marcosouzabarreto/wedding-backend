package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRSVPRoutes(r *gin.Engine, db *gorm.DB) {
	rsvp := r.Group("/rsvps")
	rsvpService := services.NewRSVPService(db)
	rsvpHandler := handlers.NewRSVPHandler(rsvpService)

	{
		rsvp.POST("/", rsvpHandler.Create)
	}

}
