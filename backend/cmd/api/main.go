package main

import (
	"log"

	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/config"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no se encontró archivo .env, usando variables de entorno del sistema")
	}

	cfg := config.LoadConfig()

	db, err := database.ConnectDatabase(cfg.GetDatabaseDSN())

	if err != nil {
		log.Fatalf("error conectando a la base de datos: %v", err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("error en migraciones: %v", err)
	}

	// // Create a Gin router with default middleware (logger and recovery)
	// r := gin.Default()

	// // Define a simple GET endpoint
	// r.GET("/ping", func(c *gin.Context) {
	// 	// Return JSON response
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// // Start server on port 8080 (default)
	// // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	// r.Run()
}
