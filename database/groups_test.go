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

func TestDeleteGroup(t *testing.T) {
	t.Run("Should delete the group (no existence check)", func(t *testing.T) {

		if db, err = GetConnection(); err != nil {
			panic("Pre test error")
		}

		mockWriter := httptest.NewRecorder()

		var groupToDelete Group
		body := GroupBody{
			Name: "My deleted group",
		}

		groupToDelete, err = CreateGroup(body, mockWriter)
		if err != nil {
			panic("Error in test")
		}

		groupId := groupToDelete.ID

		err = DeleteGroup(groupId, mockWriter)

		if err != nil {
			t.Errorf("Got an error but did not expect one : %s", err.Error())
		}

		groupWhichShouldHaveBeenDeleted := Group{}

		result := db.Find(&groupWhichShouldHaveBeenDeleted, "id = ? and deleted_at is not null", groupId)

		if result.RowsAffected != 0 {
			t.Errorf("The entity should have been deleted but it was not")
		}

	})
}
