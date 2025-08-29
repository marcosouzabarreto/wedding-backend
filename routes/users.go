package routes

import (
	"wedding-backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/users", handlers.CreateUser(db))
}
