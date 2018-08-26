package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lakshanwd/muve-go/go-crud/auth"
	"github.com/lakshanwd/muve-go/go-crud/dao"
	"github.com/lakshanwd/muve-go/go-crud/repo"
)

//Authenticate - handle authenticate requests
func Authenticate(c *gin.Context) {
	var userCredentials dao.UserCredentials
	if err := c.ShouldBindJSON(&userCredentials); err == nil {
		if user := repo.Authenticate(&userCredentials); user != nil {
			if token, err := auth.GetAuthToken(user); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": true, "token": token})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"status": false})
		}

	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
