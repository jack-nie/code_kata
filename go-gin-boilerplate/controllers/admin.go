package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/forms"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/models"
)

type AdminController struct{}

func (a AdminController) Login(c *gin.Context) {
	var adminLogin forms.AdminLogin
	if err := c.ShouldBind(&adminLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		var admin models.Admin
		if result, err := admin.Login(adminLogin); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": result})
		}
	}
}

func (a AdminController) Create(c *gin.Context) {
	var adminLogin forms.AdminLogin
	c.ShouldBind(&adminLogin)
	var admin models.Admin
	admin.Create(adminLogin)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
