package http

import "github.com/gin-gonic/gin"

func RegisterProtectedRoutes(rg *gin.RouterGroup, handler *PatientHandler) {
	patients := rg.Group("/patients")
	{
		patients.POST("", handler.CreatePatient)
		patients.PUT("/:id", handler.UpdatePatient)
		patients.PATCH("/:id", handler.PatchPatient)
		patients.GET("", handler.ListPatients)
		patients.GET("/:id", handler.GetPatient)
	}
}

// func RegisterPublicRoutes(rg *gin.RouterGroup, handler *PatientHandler) {
// 	patients := rg.Group("/patients")
// 	{
// 		patients.GET("", handler.ListPatients)
// 		patients.GET("/:id", handler.GetPatient)
// 	}
// }
