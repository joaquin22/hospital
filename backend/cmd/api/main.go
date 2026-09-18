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

	doctorApp "github.com/joaquin22/hospital-api/internal/doctors/application"
	doctorInfra "github.com/joaquin22/hospital-api/internal/doctors/infrastructure"
	doctorHTTP "github.com/joaquin22/hospital-api/internal/doctors/infrastructure/http"
	doctorPersistence "github.com/joaquin22/hospital-api/internal/doctors/infrastructure/persistence"

	specialtyApp "github.com/joaquin22/hospital-api/internal/specialty/application"
	specialtyHTTP "github.com/joaquin22/hospital-api/internal/specialty/infrastructure/http"
	specialityPersistence "github.com/joaquin22/hospital-api/internal/specialty/infrastructure/persistence"

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
	getUserUC := userApp.NewGetUserUseCase(userRepo)
	updateUserUC := userApp.NewUpdateUserUseCase(userRepo, passwordHasher)
	patchUserUC := userApp.NewPatchUserUseCase(userRepo, passwordHasher)
	userHandler := userHTTP.NewUserHandler(registerUserUC, loginUserUC, listUsersUC, getUserUC, updateUserUC, patchUserUC)

	// --- Wiring: Patients ---

	patientRepo := patientPersistence.NewGormPatientRepository(db)
	createPatientUC := patientApp.NewCreatePatientUseCase(patientRepo)
	listPatientsUC := patientApp.NewListPatientsUseCase(patientRepo)
	getPatientUC := patientApp.NewGetPatientUseCase(patientRepo)
	updatePatientUC := patientApp.NewUpdatePatientUseCase(patientRepo)
	patchPatientUC := patientApp.NewPatchPatientUseCase(patientRepo)
	patientHandler := patientHTTP.NewPatientHandler(createPatientUC, listPatientsUC, getPatientUC, updatePatientUC, patchPatientUC)

	// --- Wiring: Specialities ---

	specialityRepo := specialityPersistence.NewGormSpecialityRepository(db)
	createSpecialityUC := specialtyApp.NewCreateSpecilityUseCase(specialityRepo)
	getSpecialityUC := specialtyApp.NewGetSpecialityUseCase(specialityRepo)
	listSpecialitiesUC := specialtyApp.NewListSpecialitiesUseCase(specialityRepo)
	updateSpecialityUC := specialtyApp.NewUpdateSpecialityUseCase(specialityRepo)
	patchSpecialityUC := specialtyApp.NewPatchSpecialityUseCase(specialityRepo)
	specialityHandler := specialtyHTTP.NewSpecialityHandler(createSpecialityUC, updateSpecialityUC, patchSpecialityUC, listSpecialitiesUC, getSpecialityUC)

	// --- Wiring: Doctor ---

	// --- Register Doctor Routes ---

	doctorRepo := doctorPersistence.NewGormDoctorRepository(db)
	listDoctorUC := doctorApp.NewListDoctorUseCase(doctorRepo, getUserUC)
	getDoctorUC := doctorApp.NewGetDoctorUseCase(doctorRepo, getUserUC)
	updateDoctorUC := doctorInfra.NewTransactionalUpdateDoctor(db, passwordHasher)
	patchDoctorUC := doctorApp.NewPatchDoctorUseCase(doctorRepo, getUserUC)
	registerDoctorUC := doctorInfra.NewTransactionalRegisterDoctor(db, passwordHasher)

	doctorHandler := doctorHTTP.NewDoctorHandler(listDoctorUC, getDoctorUC, updateDoctorUC, patchDoctorUC, registerDoctorUC)

	// --- HTTP server ---

	router := gin.New()
	router.Use(middleware.Logger(), middleware.CORS(), middleware.Recovery())
	router.NoRoute(middleware.NotFound())
	router.NoMethod(middleware.MethodNotAllowed())
	v1 := router.Group("/api")

	// --- Public routes ---
	// patientHTTP.RegisterPublicRoutes(v1, patientHandler)

	// --- Protected routes ---
	protected := v1.Group("")
	protected.Use(middleware.JWTAuth(cfg.JWTSecret))
	{
		admin := protected.Group("")
		admin.Use(middleware.RequireRole("admin"))
		userHTTP.RegisterRoutes(v1, userHandler)
		patientHTTP.RegisterProtectedRoutes(admin, patientHandler)
		specialtyHTTP.RegisterRoutes(admin, specialityHandler)
		doctorHTTP.RegisterDoctorRoutes(admin, doctorHandler)
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
