package controller

import (
	"database"
	"fmt"
	"html/template"
	"model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Api_Login(c *gin.Context) {

	account := c.PostForm("uname")
	password := c.PostForm("psw")
	fmt.Println(account)
	fmt.Println(password)
	if account != "s9443112" || password != "game54176868" {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	c.JSON(200, gin.H{"description": "Post Data Success"})

}
func Api_Post(c *gin.Context) {
	ImgURL := ""
	ann_title := c.PostForm("ann_title")
	ann_time, _ := time.Parse("0000-00-00", c.PostForm("ann_time"))
	ann_content := c.PostForm("ann_content")

	Vice_title := ann_title[strings.IndexAny(ann_title, "-")+1 : len(ann_title)]
	ann_title = ann_title[0 : strings.IndexAny(ann_title, "-")-1]
	fmt.Println(ann_content)

	if strings.ContainsAny(ann_content, "<img") {
		fmt.Println(strings.Index(ann_content, "<img"))
		fmt.Println(strings.Index(ann_content, "/>"))
		ImgURL = ann_content[strings.Index(ann_content, "<img") : strings.Index(ann_content, "/></p>")+2]
		fmt.Println("ImgURL:" + ImgURL)
	}
	data := model.DataPreview{Title: ann_title, Vice_title: Vice_title, PublicTime: ann_time, User: "Doghow", Content: template.HTML(ann_content), ImgURL: template.HTML(ImgURL)}
	db := database.New()
	db.Preview.CreateOne(&data)
	c.JSON(200, gin.H{"description": "Post Data Success"})

}
