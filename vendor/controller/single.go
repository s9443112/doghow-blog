package controller

import (
	"database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	url := c.Param("url")
	db := database.New()
	preview := db.Preview.QueryOne(url)
	fmt.Print(preview)
	c.HTML(http.StatusOK, "main/info", gin.H{
		"preview": preview,
	})
}
