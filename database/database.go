package database

import (
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	Url  string
	Name string
}

type Group struct {
	gorm.Model
	Name         string
	FeedsInGroup []Feed `gorm:"many2many:group_feed;"`
}

func GetConnection() (*gorm.DB, error) {

	dsn := "user:password@tcp(127.0.0.1:3309)/feeds_database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection failed")
		return nil, errors.New("Connection failed")
	}

	return db, nil
}

func MakeMigration(db *gorm.DB) {
	db.AutoMigrate(&Feed{}, &Group{})
}
