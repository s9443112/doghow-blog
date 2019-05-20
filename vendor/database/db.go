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
	Story 		*Story
	Information 	*Information
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

	story := &Story{
		db: db,
	}

	info := &Information{
		db:db,
	}

	return &Service{
		db:			db,
		Story: 		story,
		Information: info,
	}
}

func (s *Service) InitDatabase() {
	s.db.AutoMigrate(&model.DataStory{})
	s.db.AutoMigrate(&model.DateInfo{})
}

func (s *Service) CloseDBConnect() {
	defer s.db.Close()
}
