package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestListGroups(t *testing.T) {
	t.Run("Should return a list", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/posts", nil)

		response := httptest.NewRecorder()

		ListGroups(response, req)

		if reflect.TypeOf(response.Body).Name() != "FeedList" {
			t.Errorf("Type error, expected %s got %s", "FeedList", reflect.TypeOf(response.Body).Name())
		}
	})
}
