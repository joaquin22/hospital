package http

import "github.com/gin-gonic/gin"

func RegisterProtectedRoutes(rg *gin.RouterGroup, handler *PatientHandler) {
	patients := rg.Group("/patients")
	{
		patients.POST("", handler.CreatePatient)
	}
}

func RegisterPublicRoutes(rg *gin.RouterGroup, handler *PatientHandler) {
	patients := rg.Group("/patients")
	{
		patients.GET("", handler.ListPatients)
	}
}
