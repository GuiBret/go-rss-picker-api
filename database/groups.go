package database

import (
	"net/http"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name         string `json:"name"`
	FeedsInGroup []Feed `gorm:"many2many:group_feed;"`
}

type GroupBody struct {
	Name string `json:"name"`
}

func CreateGroup(body GroupBody, w http.ResponseWriter) (Group, error) {

	var result Group

	if db, err = GetConnection(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return result, err
	}

	result = Group{
		Name:         body.Name,
		FeedsInGroup: []Feed{},
	}

	db.Create(&result)

	return result, nil
}
