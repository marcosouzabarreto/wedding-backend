package main

import (
	"fmt"
	"log"
	"wedding-backend/db"
	"wedding-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}

	router.Use(cors.Default())
	routes.SetupGuestRoutes(router)
	routes.SetupFamilyRoutes(router)

	fmt.Println("123")
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)

	}
}
