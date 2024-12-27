package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/auth"
	"rest.com/pkg/repository"
)

func Registration(c *gin.Context) {
	var inputForm auth.Credentials

	if err := c.BindJSON(&inputForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "input error"})
		return
	}

	if _, err := repository.GetUserByLogin(inputForm.Username); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage": "user already exists"})
		return
	}
	repository.AddUser(inputForm.Username, inputForm.Password)

	c.JSON(http.StatusCreated, gin.H{"message": "account created"})
}