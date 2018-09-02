package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type DataPreview struct {
	gorm.Model
	URL        string `type:varchar(100);gorm:"unique"`
	Title      string
	PublicTime time.Time
	Grade      string
	Tag        string
	PImg       string
	Content    string
	ImgJSON    string
}
