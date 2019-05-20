package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AdminIndex view
func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index", gin.H{
		"title": "管理界面",
	})
}

//AdminPreviewUpload route
func AdminPreviewUpload(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Cant Read Form : %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Cant Read Form : %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("Done , %d file", len(files)))

}
