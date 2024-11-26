package main

import (
	"webFinder/db"
	"webFinder/handlers"
	"webFinder/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitDB()

	// Initialize services
	authService := services.NewAuthService()
	searchService := services.NewSearchService()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	searchHandler := handlers.NewSearchHandler(searchService)

	// Setup Gin router
	r := gin.Default()

	// Auth routes
	r.POST("/auth/login", authHandler.Login)

	// Search routes
	r.POST("/trigger_script", searchHandler.TriggerScript)
	r.GET("/get_results/:taskID", searchHandler.GetResults)

	// Start server
	r.Run(":8080")
}
