package main

import (
	"LmsSystem/database"
	"LmsSystem/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus" // Импортируем logrus
	"log"
	"os"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	r := gin.Default()

	// Initialize the logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // Optional: Use JSON format
	logger.SetOutput(os.Stdout)

	router.SetupRoutes(db, r, logger) // Pass the logger

	fmt.Println("Registered routes:")
	for _, rt := range r.Routes() {
		fmt.Printf("%-6s %s\n", rt.Method, rt.Path)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)
	r.Run(":" + port)
}
