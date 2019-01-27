package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/forms"
)

type PostsController struct{}

func (p PostsController) Create(c *gin.Context) {
	var postContent forms.PostContent
	if err := c.ShouldBind(&postContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
