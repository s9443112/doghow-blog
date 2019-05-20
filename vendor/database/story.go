package database

import (
	"model"

	"github.com/jinzhu/gorm"
)

type Story struct {
	db *gorm.DB
}

func (i *Story) QueryAll() *[]model.DataStory {
	var story []model.DataStory
	i.db.Find(&story)
	return &story
}
