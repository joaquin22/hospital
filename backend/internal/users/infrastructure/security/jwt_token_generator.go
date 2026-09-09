package security

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joaquin22/hospital-api/internal/users/domain"
)

// Claims son los datos propios que viajan dentro del JWT, además de los
// estándar (exp, iat). "sub" (el user ID) y "role" son lo mínimo que
// necesita el middleware de autorización para proteger rutas. "sub" viaja
// como string aunque el ID interno sea uint, porque es la convención
// estándar del claim "sub" en JWT (RFC 7519).
type Claims struct {
	UserID string `json:"sub"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// JWTTokenGenerator implementa domain.TokenGenerator. Es la única pieza
// de este contexto que sabe que existe la librería JWT.
type JWTTokenGenerator struct {
	secret        []byte
	expiryMinutes int
}

func NewJWTTokenGenerator(secret string, expiryMinutes int) *JWTTokenGenerator {
	return &JWTTokenGenerator{secret: []byte(secret), expiryMinutes: expiryMinutes}
}

func (g *JWTTokenGenerator) GenerateToken(userID uint, role domain.Role) (string, int, error) {
	expiresAt := time.Now().Add(time.Duration(g.expiryMinutes) * time.Minute)

	claims := Claims{
		UserID: strconv.FormatUint(uint64(userID), 10),
		Role:   string(role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(g.secret)
	if err != nil {
		return "", 0, err
	}

	return signed, g.expiryMinutes * 60, nil
}

// ParseAndValidate se usa desde el middleware HTTP (shared/infrastructure/middleware)
// para verificar la firma y expiración de un token recibido en un request.
func ParseAndValidate(tokenString, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
