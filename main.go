package main

import (
	"controller"
	"database"
	"fmt"
	"os"
	"middleware"
	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--runserver" {
			//fmt.Println("OK Server")
			gin.SetMode(gin.ReleaseMode)
			route := gin.Default()
			route.Static("/static", "static")

			//Setting render Template
			render := eztemplate.New()
			render.TemplatesDir = "views/"
			render.Ext = ".tmpl"
			render.Debug = true

			route.HTMLRender = render.Init()
			route.GET("/", controller.Index)
			route.POST("/api",controller.AddInfo)

			route.GET("/oauth/callback",middleware.AuthGithubCallback(),controller.OauthCallback)
			
			adminGroup := route.Group("/admin")
			adminGroup.Use(middleware.AuthRequired())
			{
				adminGroup.GET("/", controller.AdminIndex)
			}

			route.Run()
		} else if os.Args[1] == "--dbinit" {
			dbinit := database.New()
			dbinit.InitDatabase()
			dbinit.CloseDBConnect()
		} else {
			fmt.Println("Please use:")
			fmt.Println("--runserver :run gin server.")
			fmt.Println("--dbinit : run mysql init.")
			os.Exit(3)
		}
	} else {
		fmt.Println("Please use:")
		fmt.Println("--runserver :run gin server.")
		fmt.Println("--dbinit : run mysql init.")
		os.Exit(3)
	}

}
