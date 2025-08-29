package main

import (
	"fmt"
	"log"
	"os"
	"wedding-backend/db"
	"wedding-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	db, err := db.InitDB()

	if err != nil {
		log.Fatalf("Failed to init DB")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.SetupGuestRoutes(router, db)
	routes.SetupFamilyRoutes(router, db)
	routes.SetupRSVPRoutes(router, db)
	routes.SetupGiftRoutes(router, db)
	routes.PaymentRoutes(router, db)

	port := os.Getenv("PORT")

	address := fmt.Sprintf("localhost:%s", port)
	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		fmt.Println("Server started at address: ")
	}
}
