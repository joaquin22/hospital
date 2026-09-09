package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/config"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/database"
	"github.com/joho/godotenv"

	userApp "github.com/joaquin22/hospital-api/internal/users/application"
	userHTTP "github.com/joaquin22/hospital-api/internal/users/infrastructure/http"
	userPersistence "github.com/joaquin22/hospital-api/internal/users/infrastructure/persistence"
	userSecurity "github.com/joaquin22/hospital-api/internal/users/infrastructure/security"
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

	userRepo := userPersistence.NewGormUserRepository(db)
	passwordHasher := userSecurity.NewBcryptHasher()
	tokenManager := userSecurity.NewJWTTokenGenerator(cfg.JWTSecret, cfg.JWTExpiryMinutes)

	registerUserUC := userApp.NewRegisterUserUseCase(userRepo, passwordHasher)
	loginUserUC := userApp.NewLoginUserUseCase(userRepo, passwordHasher, tokenManager)
	userHandler := userHTTP.NewUserHandler(registerUserUC, loginUserUC)
	// // Create a Gin router with default middleware (logger and recovery)
	router := gin.New()

	userHTTP.RegisterRoutes(router, userHandler)

	// // Define a simple GET endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("servidor escuchando en :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("error iniciando servidor: %v", err)
	}
}
