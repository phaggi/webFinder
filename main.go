package main

import (
	"net/http"
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

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Load HTML templates from the "static" directory
	r.LoadHTMLGlob("static/*")

	// Auth routes
	r.POST("/auth/login", authHandler.Login)

	// Search routes
	r.POST("/trigger_script", searchHandler.TriggerScript)
	r.GET("/get_results/:taskID", searchHandler.GetResults)

	// Home route
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start server
	r.Run(":8080")
}
