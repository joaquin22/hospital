package http

import "github.com/gin-gonic/gin"

func RegisterDoctorRoutes(rg *gin.RouterGroup, handler *DoctorHandler) {
	doctorGroup := rg.Group("/doctors")
	{
		doctorGroup.POST("", handler.RegisterDoctor)
		doctorGroup.PUT("/:id", handler.UpdateDoctor)
		doctorGroup.PATCH("/:id", handler.PatchDoctor)
		doctorGroup.GET("", handler.ListDoctors)
		doctorGroup.GET("/:id", handler.GetDoctor)
	}
}
