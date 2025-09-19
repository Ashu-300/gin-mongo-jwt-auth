package middleware

import (
	"gin-jwt-auth/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}

		token := parts[1]

		user , err := jwt.VerifyToken(token)
		
		if err!= nil || user == nil {
			c.JSON(401, gin.H{"message": "wrong token", "error": err.Error()})
			c.Abort()  // stop chain
			return
		}

		c.Set("userInfo" , user) 
		c.Next()
	}
}