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

	var group Group
	db, err := GetConnection()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return group, err
	}

	group = Group{
		Name:         body.Name,
		FeedsInGroup: []Feed{},
	}

	db.Create(&group)

	return group, nil
}
