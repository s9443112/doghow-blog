package model

import (
	"html/template"

	"github.com/jinzhu/gorm"
)

type DataProfile struct {
	gorm.Model
	Icom    string
	Name    string `type:varchar(100);gorm:"unique"`
	Job     string
	Content template.HTML
}
