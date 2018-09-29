package model

import (
	"html/template"
	"time"

	"github.com/jinzhu/gorm"
)

type DataPreview struct {
	gorm.Model
	Title      string
	Vice_title string
	PublicTime time.Time
	User       string
	ImgURL     template.HTML
	Content    template.HTML
}
