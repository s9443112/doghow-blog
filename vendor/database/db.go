package database

import (
	"fmt"
	"local_modules/configloader"
	"model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Service struct {
	db          *gorm.DB
	Information *Information
}

func New() *Service {
	c := configloader.New("config.yaml")
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.DB.User, c.DB.Password, c.DB.Host, c.DB.DBName)
	fmt.Println(connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic("failed to connect database")
	}
	db.SingularTable(true) //不會幫你把table name +s

	info := &Information{
		db: db,
	}
	return &Service{
		db:          db,
		Information: info,
	}
}

func (s *Service) InitDatabase() {
	s.db.AutoMigrate(&model.DataInfo{})
	s.db.AutoMigrate(&model.DataPreview{})
	s.db.AutoMigrate(&model.DataProfile{})
}

func (s *Service) CloseDBConnect() {
	defer s.db.Close()
}
