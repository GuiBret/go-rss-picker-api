package database

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	Url  string
	Name string
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

func HandleConnectionError(err error, w http.ResponseWriter) {
	fmt.Fprintf(w, "DB connection error : %s", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}
