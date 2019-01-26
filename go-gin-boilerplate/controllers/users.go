package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/forms"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/models"
	. "github.com/jack-nie/code_kata/go-gin-boilerplate/models"
)

type UsersController struct{}

func (u UsersController) Signup(c *gin.Context) {
	var userSignup forms.UserSignup
	var user models.User

	if err := c.ShouldBind(&userSignup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		user.Signup(userSignup)
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (u UsersController) Show(c *gin.Context) {
	var user models.User
	id := c.Query("id")
	DB.First(&user, id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
