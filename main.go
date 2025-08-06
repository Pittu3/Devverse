package main

import (
	"devverse-backend/config"
	"devverse-backend/models"
	"devverse-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Admin{})

	router := gin.Default()

	// ✅ Allow cross-origin requests
	router.Use(cors.Default())

	routes.SetupRoutes(router)

	router.Run(":5000")
}
