package util

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint   `json:"uid"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(secret string, uid uint, role, email, name string) (string, error) {
	claims := JWTClaims{
		UserID: uid,
		Role:   role,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GetUserClaims helper untuk mengambil claims opsional (tanpa verifikasi ulang)
func GetUserClaims(c *fiber.Ctx) JWTClaims {
	auth := c.Get("Authorization")
	if len(auth) > 7 {
		tokenStr := auth[7:]
		token, _ := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) { return []byte(""), nil })
		if token != nil {
			if cl, ok := token.Claims.(*JWTClaims); ok {
				return *cl
			}
		}
	}
	return JWTClaims{}
}
