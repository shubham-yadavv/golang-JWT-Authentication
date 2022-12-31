package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {

	tokenString, err := c.Cookie("authtoken")

	if err != nil {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("user", claims["user"])
		c.Next()
	}

}
