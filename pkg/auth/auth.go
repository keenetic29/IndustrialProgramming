package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	ADMIN_ROLE = "Admin"
	USER_ROLE  = "User"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username 	string `json:"username"`
	AccessLevel	string `json:"accesslevel"`
	jwt.StandardClaims
}

func GenerateToken(username string, accesslevel string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: 		username,
		AccessLevel:    accesslevel,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.AccessLevel)

		c.Next()
	}
}

func AdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		if role, exists := c.Get("role"); !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return

		} else if role != ADMIN_ROLE {
			c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to access this resource"})
			c.Abort()
		}

		c.Next()
	}
}