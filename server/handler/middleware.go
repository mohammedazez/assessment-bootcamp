package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/auth"
	"server/users"
	"strconv"
)

func Middleware(userService users.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		token, err := authService.ValidateToken(header)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		userID := int(claim["user_id"].(float64))

		userIDstr := strconv.Itoa(userID)

		user, err := userService.GetUserByIDService(userIDstr)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}
		if user.ID == 0 {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		c.Set("currentUser", user)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
