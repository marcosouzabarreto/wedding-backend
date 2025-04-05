package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
)

func SetupFamilyRoutes(r *gin.Engine) {
	family := r.Group("/families")
	familyService := services.NewFamilyService()
	familyHandlers := handlers.NewFamilyHandlers(familyService)

	{
		family.GET("/", familyHandlers.GetAll)
		family.POST("/", familyHandlers.Create)
		family.GET("/:id", familyHandlers.GetById)
	}

}

