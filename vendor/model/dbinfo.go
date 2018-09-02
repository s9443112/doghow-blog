package model

import "github.com/jinzhu/gorm"

type DataInfo struct {
	gorm.Model
	URL     string `type:varchar(100);gorm:"unique"`
	Title   string
	Content string
}
