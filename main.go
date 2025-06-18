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
	routes.SetupRSVPRoutes(router, db)

	port := "8080"
	address := fmt.Sprintf("localhost:%s", port)
	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		fmt.Println("Server started at address: ")
	}
}
