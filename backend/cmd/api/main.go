package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/config"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/database"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/middleware"
	"github.com/joho/godotenv"

	patientApp "github.com/joaquin22/hospital-api/internal/patients/application"
	patientHTTP "github.com/joaquin22/hospital-api/internal/patients/infrastructure/http"
	patientPersistence "github.com/joaquin22/hospital-api/internal/patients/infrastructure/persistence"

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

	// --- Wiring: User (auth) ---
	userRepo := userPersistence.NewGormUserRepository(db)
	passwordHasher := userSecurity.NewBcryptHasher()
	tokenManager := userSecurity.NewJWTTokenGenerator(cfg.JWTSecret, cfg.JWTExpiryMinutes)
	registerUserUC := userApp.NewRegisterUserUseCase(userRepo, passwordHasher)
	loginUserUC := userApp.NewLoginUserUseCase(userRepo, passwordHasher, tokenManager)
	listUsersUC := userApp.NewListUsersUseCase(userRepo)
	userHandler := userHTTP.NewUserHandler(registerUserUC, loginUserUC, listUsersUC)

	// --- Wiring: Patients ---

	patientRepo := patientPersistence.NewGormPatientRepository(db)
	createPatientUC := patientApp.NewCreatePatientUseCase(patientRepo)
	listPatientsUC := patientApp.NewListPatientsUseCase(patientRepo)
	getPatientUC := patientApp.NewGetPatientUseCase(patientRepo)
	updatePatientUC := patientApp.NewUpdatePatientUseCase(patientRepo)
	patchPatientUC := patientApp.NewPatchPatientUseCase(patientRepo)
	patientHandler := patientHTTP.NewPatientHandler(createPatientUC, listPatientsUC, getPatientUC, updatePatientUC, patchPatientUC)

	// --- HTTP server ---

	router := gin.New()
	router.Use(middleware.Logger(), middleware.CORS(), gin.Recovery())
	v1 := router.Group("/api")

	// --- Public routes ---
	userHTTP.RegisterRoutes(v1, userHandler)
	// patientHTTP.RegisterPublicRoutes(v1, patientHandler)

	// --- Protected routes ---
	protected := v1.Group("")
	protected.Use(middleware.JWTAuth(cfg.JWTSecret))
	{
		admin := protected.Group("")
		admin.Use(middleware.RequireRole("admin"))
		patientHTTP.RegisterProtectedRoutes(admin, patientHandler)
	}

	// // Define a simple GET endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("servidor escuchando en :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("error iniciando servidor: %v", err)
	}
}
