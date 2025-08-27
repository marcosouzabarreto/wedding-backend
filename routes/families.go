package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupFamilyRoutes(r *gin.Engine, db *gorm.DB) {
	family := r.Group("/families")
	familyService := services.NewFamilyService(db)
	familyHandlers := handlers.NewFamilyHandlers(familyService)

	{
		family.GET("/", familyHandlers.GetAll)
		family.POST("/", familyHandlers.Create)
		family.GET("/:id", familyHandlers.GetById)
		family.GET("/token/:token", familyHandlers.GetByToken)
	}

}

