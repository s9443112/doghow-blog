package database

import (
	"model"

	"github.com/jinzhu/gorm"
)

type Preview struct {
	db *gorm.DB
}

func (p *Preview) QueryLimit(count int) *[]model.DataPreview {
	var preview []model.DataPreview
	p.db.Limit(count).Find(&preview)
	return &preview
}

func (p *Preview) QueryOne(url string) *model.DataPreview {
	var preview model.DataPreview
	p.db.Where(url).First(&preview)
	return &preview
}

func (p *Preview) CreateOne(dPreview *model.DataPreview) {
	p.db.Create(&dPreview)
}
