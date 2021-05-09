package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"rss-picker-api/database"
	"strings"
	"testing"
)

func TestListGroups(t *testing.T) {
	t.Run("Should return a list", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "http://localhost:4005/groups", nil)

		v := GroupList{}
		response := httptest.NewRecorder()

		ListGroups(response, req)

		body, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal(body, &v)

		if err != nil {
			t.Errorf("Type error : %s", err)
		}
	})
}

func TestListFeeds(t *testing.T) {
	t.Run("Should return a list of feeds", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://localhost:4005/feeds", nil)

		v := FeedList{}

		response := httptest.NewRecorder()

		ListFeeds(response, req)

		body, _ := ioutil.ReadAll(response.Body)

		err := json.Unmarshal(body, &v)

		if err != nil {
			t.Errorf("Type error : %s", err)
		}

	})
}

func TestAddGroup(t *testing.T) {
	t.Run("Should return an error since the body is invalid", func(t *testing.T) {
		body := strings.NewReader(``)

		req, _ := http.NewRequest("POST", "http://localhost:4005/groups", body)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(AddGroup)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Expected HTTP 400 but got %d", status)
		}

	})
	t.Run("Should return an error since name is not passed", func(t *testing.T) {
		body := strings.NewReader(`{}`)

		req, _ := http.NewRequest("POST", "http://localhost:4005/groups", body)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(AddGroup)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Expected HTTP 400 but got %d", status)
		}

	})

	t.Run("Should return HTTP Created since all fields are passed", func(t *testing.T) {
		body := strings.NewReader(`{"name": "CNN"}`)

		req, _ := http.NewRequest("POST", "http://localhost:4005/groups", body)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(AddGroup)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Expected HTTP 201 but got %d", status)
		}

	})
}

func TestAddFeed(t *testing.T) {
	t.Run("Should return an error since the body is invalid", func(t *testing.T) {
		body := strings.NewReader(``)

		req, _ := http.NewRequest("POST", "http://localhost:4005/feeds", body)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(AddFeed)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Status error, expected %d got %d", http.StatusBadRequest, status)
		}
	})

	t.Run("Should return an error since the body is invalid", func(t *testing.T) {
		body := strings.NewReader(`{}`)

		req, _ := http.NewRequest("POST", "http://localhost:4005/feeds", body)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(AddFeed)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Status error, expected %d got %d", http.StatusBadRequest, status)
		}
	})

	t.Run("Should return HTTP created since everything is valid", func(t *testing.T) {
		body := strings.NewReader(`{"name": "My feed", "url": "https://cnn.com/feed"}`)

		req, _ := http.NewRequest("POST", "http://localhost:4005/feeds", body)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(AddFeed)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Status error, expected %d got %d", http.StatusCreated, status)
		}
	})
}

func TestDeleteFeed(t *testing.T) {
	t.Run("Should return an error since the argument passed is not a number", func(t *testing.T) {

		// First, we add a new feed, and we will use the created ID to delete it
		feed := database.FeedBody{
			Name: "abc",
			Url:  "https://def.com/url",
		}

		mockRR := httptest.NewRecorder()

		id, err := database.CreateFeed(feed, mockRR)

		if err != nil {
			panic("Test error")
		}

		req, _ := http.NewRequest("DELETE", "http://localhost:4005/feeds/"+fmt.Sprint(id), nil)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(DeleteFeed)

		handler.ServeHTTP(rr, req)

		// The feed should not exist anymore

	})
}
