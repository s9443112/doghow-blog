package model

import "github.com/jinzhu/gorm"

type DataStory struct {
	gorm.Model
	Title   string
	Content string
	Img 	string
}
