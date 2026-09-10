package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/users/infrastructure/security"
)

// JWTAuth es un middleware TÉCNICO: verifica que el header Authorization
// traiga un "Bearer <token>" válido y no expirado, y deja el userID/role
// disponibles en el contexto de Gin para que los handlers los usen si
// necesitan autorización más fina (ej. "solo ADMIN puede...").
//
// Vive en shared/ (y no dentro de internal/user/) porque lo usan TODOS los
// contextos que necesiten proteger rutas (reservation, room), no solo
// user — es infraestructura transversal, no una regla de negocio de
// ningún contexto en particular. Solo importa el paquete `security` de
// user/infrastructure para parsear el token, nunca el dominio de user.
func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "falta el header Authorization: Bearer <token>"})
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		claims, err := security.ParseAndValidate(tokenString, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido o expirado"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

// RequireRole se encadena DESPUÉS de JWTAuth y restringe el endpoint a
// uno o más roles específicos (ej. solo ADMIN puede crear habitaciones).
func RequireRole(allowed ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("user_role")
		for _, r := range allowed {
			if role == r {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no tienes permisos para esta acción"})
	}
}
