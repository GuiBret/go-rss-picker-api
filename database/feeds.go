package database

import (
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type FeedBody struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

var err error
var db *gorm.DB

func CreateFeed(body FeedBody, w http.ResponseWriter) (uint, error) {

	if db, err = GetConnection(); err != nil {
		HandleConnectionError(err, w)
		return 0, err
	}

	feed := Feed{
		Url:  body.Url,
		Name: body.Name,
	}

	db.Create(&feed)
	return feed.ID, nil
}

func GetFeed(id uint, w http.ResponseWriter) (Feed, error) {

	if db, err = GetConnection(); err != nil {
		HandleConnectionError(err, w)
		return Feed{}, err

	}

	feed := Feed{}

	result := db.Find(&feed, "id = ?", id)

	if result.RowsAffected == 0 {

		w.WriteHeader(http.StatusNotFound)
		return feed, errors.New(fmt.Sprintf("No record found for %d", id))
	}

	return feed, nil
}

func ListFeeds(w http.ResponseWriter) ([]Feed, error) {
	var feeds []Feed

	if db, err = GetConnection(); err != nil {
		HandleConnectionError(err, w)
		return feeds, err
	}

	db.Find(&feeds)

	return feeds, nil
}

func DeleteFeed(id uint, w http.ResponseWriter) error {

	if db, err = GetConnection(); err != nil {
		HandleConnectionError(err, w)
		return err

	}

	feed := Feed{}

	result := db.Find(&feed, "id = ?", id)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return errors.New(fmt.Sprintf("The entity with ID %d does not exist", id))
	}

	db.Delete(&Feed{}, id)

	return nil

}
