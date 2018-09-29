package controller

import (
	"database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	db := database.New()
	info := db.Preview.QueryLimit(5)
	c.HTML(http.StatusOK, "main/index", gin.H{
		"info": info,
	})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "main/login", gin.H{})
}
func Post(c *gin.Context) {
	c.HTML(http.StatusOK, "main/post", gin.H{})
}
