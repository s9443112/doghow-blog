package database

import (
	"model"

	"github.com/jinzhu/gorm"
)

type Information struct {
	db *gorm.DB
}

func (i *Information) QueryOne(url string) *[]model.DateInfo {
	var info []model.DateInfo
	i.db.Where("URL = ?", url).First(&info)
	return &info
}

func (i *Information) Add(info *model.DateInfo) {
	i.db.Create(&info)
}
