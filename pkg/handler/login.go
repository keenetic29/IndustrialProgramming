package handler

import (
	"net/http"

	"rest.com/pkg/auth"
	"rest.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        credentials  body      auth.Credentials  true  "User credentials"
// @Success      200  {object}  map[string]string  "JWT token"
// @Failure      400  {object}  map[string]string  "Invalid request body"
// @Failure      401  {object}  map[string]string  "Unauthorized"
// @Failure      404  {object}  map[string]string  "User not found"
// @Failure      500  {object}  map[string]string  "Failed to generate token"
// @Router       /login [post]
func Login(c *gin.Context) {
	var creds auth.Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	user, err := repository.GetUserByLogin(creds.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if creds.Username != user.Username || creds.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := auth.GenerateToken(creds.Username, user.AccessLevel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}