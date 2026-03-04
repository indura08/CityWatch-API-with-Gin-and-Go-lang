package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeRoles(secret string, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "Authorization header missing",
			})
			c.Abort()
			return
		}

		//check the token is valid or not
		parts := strings.Split(authHeader, "")
		if len(parts) != 2 {
			c.JSON(401, gin.H{
				"error": "Invalid Token",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{
				"error": "Invalid Token",
			})
			c.Abort()
		}

		//need to implement the other parts
	}
}
