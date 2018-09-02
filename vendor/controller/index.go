package controller

import (
	"database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	db := database.New()
	info := db.Information.QueryLimit(10)
	c.HTML(http.StatusOK, "main/index", gin.H{
		"info": info,
	})
}
