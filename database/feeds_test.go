package database

import (
	"net/http/httptest"
	"testing"
)

func TestCreateFeed(t *testing.T) {
	t.Run("Should return an ID upon creating ", func(t *testing.T) {

		feed := FeedBody{
			Name: "abc",
			Url:  "https://def.com/feeds",
		}

		mockWriter := httptest.NewRecorder()
		id, err := CreateFeed(feed, mockWriter)

		if id == 0 && err != nil {
			t.Errorf("Expected an ID but got 0")
		}
	})
}
