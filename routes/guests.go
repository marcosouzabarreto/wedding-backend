package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
)

func SetupGuestRoutes(r *gin.Engine) {
	guest := r.Group("/guests")
	guestService := services.NewGuestService()
	guestRoutes := handlers.NewGuestHandlers(guestService)
	{
		guest.GET("/", guestRoutes.GetAll)
		guest.POST("/", guestRoutes.Create)
		guest.GET("/:id", guestRoutes.GetByID)
	}
}
