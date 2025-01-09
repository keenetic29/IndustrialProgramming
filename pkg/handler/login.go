package handler

import (
	"net/http"

	"rest.com/pkg/auth"
	"rest.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

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