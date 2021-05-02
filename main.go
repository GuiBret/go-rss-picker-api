package main

import (
	"net/http"
	"rss-picker-api/database"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var myRouter mux.Router
var db *gorm.DB

type Feed struct {
	gorm.Model
	Id  int
	Url string
}

type Group struct {
	gorm.Model
	id           int
	name         string
	FeedsInGroup []Feed `gorm:"many2many:group_feed;"`
}

func main() {

	db, err := database.GetConnection()

	if err != nil {
		panic(err.Error())
	}

	database.MakeMigration(db)

	handleRequests()
}

func handleRequests() {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/groups", ListGroups).Methods("GET")
	myRouter.HandleFunc("/groups", AddGroup).Methods("POST")
	myRouter.HandleFunc("/groups", DeleteGroup).Methods("DELETE")

	myRouter.HandleFunc("/groups/{groupId}/feeds", ListFeedsInGroup).Methods("GET")
	myRouter.HandleFunc("/groups/{feedId}/feeds/{feedId}", AddFeedToGroup).Methods("POST")
	myRouter.HandleFunc("/groups/{feedId}/feeds/{feedId}", DeleteFeedFromGroup).Methods("DELETE")

}

func ListGroups(w http.ResponseWriter, r *http.Request) {
}

func AddGroup(w http.ResponseWriter, r *http.Request) {

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}

func ListFeedsInGroup(w http.ResponseWriter, r *http.Request) {

}

func AddFeedToGroup(w http.ResponseWriter, r *http.Request) {

}

func DeleteFeedFromGroup(w http.ResponseWriter, r *http.Request) {

}
