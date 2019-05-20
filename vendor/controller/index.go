package controller

import (
	"database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	db := database.New()
	story := db.Story.QueryAll()
	c.HTML(http.StatusOK, "main/index", gin.H{
		"story": story,
	})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "main/login", gin.H{})
}
func Post(c *gin.Context) {
	c.HTML(http.StatusOK, "main/post", gin.H{})
}
