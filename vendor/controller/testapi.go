package controller

import (
	"github.com/gin-gonic/gin"
)

type Postdata struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Testapi(c *gin.Context) {
	var p Postdata
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
	}
	c.JSON(200, gin.H{
		"status":  "post",
		"message": "ok successed",
	})

}
