package controller

import (
	"github.com/gin-gonic/gin"
)

func OauthCallback(c *gin.Context){
	c.Redirect(302,"/admin")	
}