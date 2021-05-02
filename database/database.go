package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {

	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connection failed")
	}

	return db
}
