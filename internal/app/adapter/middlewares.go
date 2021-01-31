package adapter

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Dalot/goddd/internal/app/application/usecase"
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.String(http.StatusOK, "")
			return
		}
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		header := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(header, "Bearer ")
		if len(splitToken) > 1 {
			tokenStr = splitToken[1]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"status":  "error",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return usecase.JwtKey, nil
		})
		
		var user domain.User

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user.Email = claims["email"].(string)

			log.Println(user)
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  "error",
			})
			c.Abort()
			return
		}
	}
}
