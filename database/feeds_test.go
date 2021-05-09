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

func TestGetFeed(t *testing.T) {
	t.Run("Should return an error since the entity does not exist", func(t *testing.T) {
		db, err := GetConnection()

		// First, we need the max ID in the DB
		if err != nil {
			panic("Test error")
		}

		feed := Feed{}

		db.Last(&feed)

		maxId := feed.ID

		mockRR := httptest.NewRecorder()

		_, err = GetFeed(maxId+5, mockRR)

		if err == nil {

			t.Errorf("Expected an error but did not get one")
		}

	})
	t.Run("Should return the entity since it exists", func(t *testing.T) {

		// First, we need to create an entity
		feed := FeedBody{
			Name: "a",
			Url:  "https://def.com/feeds",
		}
		mockRR := httptest.NewRecorder()

		id, _ := CreateFeed(feed, mockRR)

		_, err := GetFeed(id, mockRR)

		if err != nil {

			t.Errorf("Expected no error but got %s", err.Error())
		}

	})
}

func TestDeleteFeed(t *testing.T) {
	t.Run("Should not have done anything since the element does not exist", func(t *testing.T) {

		db, err := GetConnection()
		mockRR := httptest.NewRecorder()

		// First, we need the max ID in the DB
		if err != nil {
			panic("Test error")
		}

		feed := Feed{}

		db.Last(&feed)

		maxId := feed.ID

		err = DeleteFeed(maxId+5, mockRR)

		if err == nil {
			t.Errorf("Expected an error but did not get one")
		}
	})
	t.Run("Should have deleted the entity we have just created", func(t *testing.T) {

		mockRR := httptest.NewRecorder()

		// First, we need to create an entity
		feed := FeedBody{
			Name: "a",
			Url:  "https://def.com/feeds",
		}

		id, _ := CreateFeed(feed, mockRR)

		err := DeleteFeed(id, mockRR)

		if err != nil {
			t.Errorf("Expected no error but got %s", err.Error())
		}

		_, err = GetFeed(id, mockRR)

		if err == nil {
			t.Errorf("Expected an error but did not get one")
		}

	})
}
