package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"rss-picker-api/database"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var a App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestMain(m *testing.M) {
	a.Initialize()

	code := m.Run()

	os.Exit(code)

}

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

		req, _ := http.NewRequest("POST", "/groups", body)

		rr := executeRequest(req)

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

		req, _ := http.NewRequest("DELETE", "/feeds/abcdef/", nil)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(DeleteFeed)

		handler.ServeHTTP(rr, req)

		// The request should have returned 400 since the argument is not a number

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Unexpected return code, got %d want %d", status, http.StatusBadRequest)
		}

	})

	// TODO: create entry point
	t.Run("Should return an error since the entity does not exist", func(t *testing.T) {

		db, err := database.GetConnection()
		rr := httptest.NewRecorder()

		// First, we need the max ID in the DB
		if err != nil {
			panic("Test error")
		}

		var feed database.Feed

		db.Last(&feed)

		maxId := feed.ID

		// TODO: rewrite all tests like this one
		router := mux.NewRouter()

		req, _ := http.NewRequest("DELETE", "/feeds/"+fmt.Sprint(maxId+1)+"/", nil)
		router.ServeHTTP(rr, req)

		// // The request should have returned 404 since the entity was not found
		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("Unexpected return code, got %d want %d", status, http.StatusNotFound)
		}

	})

	t.Run("Should delete the entity since it exists", func(t *testing.T) {
		feed := database.FeedBody{
			Name: "a",
			Url:  "https://def.com/feeds",
		}

		mockRR := httptest.NewRecorder()

		id, err := database.CreateFeed(feed, mockRR)

		if err != nil {
			panic("Pre test error")
		}

		req, _ := http.NewRequest("DELETE", "/feeds/"+fmt.Sprint(id), nil)

		rr := executeRequest(req)

		// The request should have returned 200 since the entity was deleted
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Unexpected return code, got %d want %d", status, http.StatusOK)
		}

	})
}
