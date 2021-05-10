package database

import (
	"net/http/httptest"
	"testing"
)

func TestCreateGroup(t *testing.T) {

	t.Run("Should succesfully create a group", func(t *testing.T) {

		var body GroupBody
		mockWriter := httptest.NewRecorder()

		body = GroupBody{
			Name: "abcdef",
		}

		group, err := CreateGroup(body, mockWriter)

		if err != nil {
			t.Errorf("Got an error but did not expect one : %s", err.Error())
		}

		if group.Name != "abcdef" || group.ID == 0 {
			t.Error("Error when creating the entity")
		}
	})

}
