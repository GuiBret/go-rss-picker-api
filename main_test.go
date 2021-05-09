package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
		body := strings.NewReader(`{}`)

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
