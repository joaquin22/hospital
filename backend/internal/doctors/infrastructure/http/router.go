package http

import "github.com/gin-gonic/gin"

func RegisterDoctorRoutes(rg *gin.RouterGroup, handler *DoctorHandler) {
	doctorGroup := rg.Group("/doctors")
	{
		doctorGroup.POST("/", handler.RegisterDoctor)
		doctorGroup.GET("/", handler.ListDoctors)
	}
}
