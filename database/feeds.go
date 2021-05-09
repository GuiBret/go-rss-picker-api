package database

import (
	"errors"
	"fmt"
	"net/http"
)

type FeedBody struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func HandleConnectionError(err error, w http.ResponseWriter) {
	fmt.Fprintf(w, "DB connection error : %s", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}

func CreateFeed(body FeedBody, w http.ResponseWriter) (uint, error) {
	db, err := GetConnection()

	if err != nil {
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
	db, err := GetConnection()

	if err != nil {
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

func DeleteFeed(id uint, w http.ResponseWriter) error {

	db, err := GetConnection()

	if err != nil {
		HandleConnectionError(err, w)
		return err

	}

	feed := Feed{}

	result := db.Find(&feed, "id = ?", id)

	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("The entity with ID %d does not exist", id))
	}

	db.Delete(&Feed{}, id)

	return nil

}
