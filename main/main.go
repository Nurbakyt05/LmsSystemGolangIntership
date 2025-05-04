package main

import (
	"LmsSystem/database"
	"LmsSystem/router"
	"fmt" // ← добавили
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	r := gin.Default()
	router.SetupRoutes(db, r)

	// ← Ниже выводим все зарегистрированные маршруты
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
