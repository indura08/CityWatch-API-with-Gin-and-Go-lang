package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeRoles(secret string, allowedRoles ...int) gin.HandlerFunc {
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
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.JSON(401, gin.H{
				"error": "Invalid Token1",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{
				"error": "Invalid Token2",
			})
			c.Abort()
			return
		}

		//need to implement the other parts
		claims := token.Claims.(jwt.MapClaims)

		rolefloat, ok := claims["role"].(float64)
		role := int(rolefloat)
		log.Print("role is :", role)

		if !ok {
			c.JSON(403, gin.H{
				"error": "Role couldnt be found in the token",
			})
			c.Abort()
			return
		}

		//check if the roles are there in the role array api pass krpu array ek me function ekt
		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		c.JSON(403, gin.H{
			"Middleware error": "Access denied",
		})
		c.Abort()
	}
}
