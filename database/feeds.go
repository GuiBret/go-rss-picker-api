package database

import (
	"fmt"
	"net/http"
)

type FeedBody struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func CreateFeed(body FeedBody, w http.ResponseWriter) (uint, error) {
	db, err := GetConnection()

	if err != nil {
		fmt.Fprintf(w, "DB connection error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return 0, err

	}
	feed := Feed{
		Url:  body.Url,
		Name: body.Name,
	}

	db.Create(&feed)
	return feed.ID, nil
}
