package routers

import (
	"github.com/gin-gonic/gin"
	"controller"
	"middleware"

	eztemplate "github.com/michelloworld/ez-gin-template"
)





func InitRouter() *gin.Engine {

	router := gin.Default()

	//Setting render Template
	render := eztemplate.New()
	render.TemplatesDir = "views/"
	render.Ext = ".tmpl"
	render.Debug = true

	router.HTMLRender = render.Init()


	router.Static("/static", "static")
	router.NoRoute(controller.NoRoute)
	router.GET("/", controller.Index)
	router.GET("/story",controller.Story)
	router.GET("/oauth/callback",middleware.AuthGithubCallback(),controller.OauthCallback)
	router.POST("/api",controller.AddInfo)


	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.AuthRequired())
	{
		adminGroup.GET("/", controller.AdminIndex)
	}
	return router
	
}
