package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
)
// 注意在struct 一定要定義大寫
// binding: required 可設定需要給值or不給
type s_data struct {
	Name       string   `json:"name" binding:"required"`
	Title 	string `json:"title"`
}
   

func AddInfo(c *gin.Context){
	fmt.Print("I'm here")
	var p s_data
	_timestamp_format := "2006-01-02 15:04:05"
	if c.BindJSON(&p) ==nil{
		
		c.JSON(200,gin.H{
			"name":p.Name,
			"title":p.Title,
		})
		
	}else{
		c.JSON(http.StatusBadRequest,gin.H{
			"errCode":4004,
			"message":"Wroung request format",
			"timestamp": time.Now().Format(_timestamp_format),
		})
	}
	
}