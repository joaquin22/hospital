package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Printf("panic recuperado: %v", recovered)
		response.ErrorResponse(c, http.StatusInternalServerError, "error interno del servidor", fmt.Errorf("%v", recovered))
	})
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ErrorResponse(c, http.StatusNotFound, "recurso no encontrado", fmt.Errorf("recurso no encontrado"))
	}
}

func MethodNotAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ErrorResponse(c, http.StatusMethodNotAllowed, "método no permitido para esta ruta", fmt.Errorf("método no permitido para esta ruta"))
	}
}
