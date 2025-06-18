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
	db, err := db.InitDB()

	if err != nil {
		log.Fatalf("Failed to init DB")
	}

	router.Use(cors.Default())
	routes.SetupGuestRoutes(router, db)
	routes.SetupFamilyRoutes(router, db)

	fmt.Println("123")
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)

	}
}
