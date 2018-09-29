package database

import (
	"model"

	"github.com/jinzhu/gorm"
)

type Information struct {
	db *gorm.DB
}

func (i *Information) QueryLimit(count int) *[]model.DataInfo {
	var info []model.DataInfo
	i.db.Limit(count).Find(&info)
	return &info
}

func (i *Information) QueryOne(url string) *model.DataInfo {
	var info model.DataInfo
	i.db.Where("URL = ?", url).First(&info)
	return &info
}
