package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *SpecialityHandler) {
	specialty := rg.Group("/specialty")
	{
		specialty.GET("", handler.ListSpecialities)
		specialty.GET("/:id", handler.GetSpeciality)
		specialty.POST("", handler.CreateSpeciality)
		specialty.PUT("/:id", handler.UpdateSpeciality)
		specialty.PATCH("/:id", handler.PatchSpeciality)
	}
}
