package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

		fmt.Printf("%s", response.Body.String())
		if err != nil {
			t.Errorf("Type error : %s", err)
		}
	})
}
