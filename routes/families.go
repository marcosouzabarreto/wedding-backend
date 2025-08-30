package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/middleware"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupFamilyRoutes(r *gin.Engine, db *gorm.DB) {
	familyService := services.NewFamilyService(db)
	familyHandlers := handlers.NewFamilyHandlers(familyService)

	familyPublic := r.Group("/families")
	{
		familyPublic.GET("/:id", familyHandlers.GetById)
		familyPublic.GET("/token/:token", familyHandlers.GetByToken)
	}

	familyProtected := r.Group("/families")
	familyProtected.Use(middleware.AuthMiddleware())
	{
		familyPublic.GET("/", familyHandlers.GetAll)
		familyProtected.POST("/", familyHandlers.Create)
	}

}

