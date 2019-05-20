package model

import (
	"github.com/jinzhu/gorm"
)

type DateInfo struct {
	gorm.Model
	Name string
	Birth string
	Email string
	Content string
	Github string `unique:gorm`
}