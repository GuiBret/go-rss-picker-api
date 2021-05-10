package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rss-picker-api/database"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var myRouter mux.Router
var db *gorm.DB

type Feed struct {
	gorm.Model
	Url  string `json:"url"`
	Name string `json:"name"`
}

type Group struct {
	gorm.Model
	Name         string `json:"name"`
	FeedsInGroup []Feed `gorm:"many2many:group_feed;"`
}

type FeedList struct {
	Feeds []Feed `json:"feeds"`
}

type GroupList struct {
	Groups []Group `json:"groups"`
}

type GroupBody struct {
	Name string `json:"name"`
}

var group Group

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {

	db, err := database.GetConnection()

	if err != nil {
		panic("Error")

	}

	a.Router = mux.NewRouter()

	a.handleRequests()
	a.DB = db
	database.MakeMigration(db)

}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":4005", a.Router))
}

func main() {
	a := App{}

	a.Initialize()

	a.Run()

}

func SetHeaderAsJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})

}

func (a *App) handleRequests() {

	// Middleware setting all return types to JSON
	a.Router.Use(SetHeaderAsJSON)

	a.Router.HandleFunc("/groups", ListGroups).Methods("GET")
	a.Router.HandleFunc("/groups", AddGroup).Methods("POST")
	a.Router.HandleFunc("/groups/{groupId:[0-9]+}", DeleteGroup).Methods("DELETE")

	a.Router.HandleFunc("/groups/{groupId}/feeds", ListFeedsInGroup).Methods("GET")
	a.Router.HandleFunc("/groups/{groupId}/feeds/{feedId}", AddFeedToGroup).Methods("POST")
	a.Router.HandleFunc("/groups/{groupId}/feeds/{feedId}", RemoveFeedFromGroup).Methods("DELETE")

	a.Router.HandleFunc("/feeds/{feedId}/", DeleteFeed).Methods("DELETE")
	a.Router.HandleFunc("/feeds", ListFeeds).Methods("GET")
	a.Router.HandleFunc("/feeds", AddFeed).Methods("POST")

}

func ListGroups(w http.ResponseWriter, r *http.Request) {

	db, err := database.GetConnection()

	if err != nil {
		panic(err.Error())
	}

	var groups []Group
	var groupResult GroupList

	db.Find(&groups)

	groupResult.Groups = groups

	json.NewEncoder(w).Encode(groupResult)

}

func AddGroup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	body := GroupBody{}
	err := decoder.Decode(&body)

	db, err := database.GetConnection()

	if err != nil {

		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	if body.Name == "" {
		http.Error(w, "Missing value 'name'", http.StatusBadRequest)
		return
	}

	group := Group{
		Name:         body.Name,
		FeedsInGroup: []Feed{},
	}

	db.Create(&group)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(nil)

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}

func ListFeeds(w http.ResponseWriter, r *http.Request) {

	db, err := database.GetConnection()

	if err != nil {
		panic(err.Error())
	}

	var groups []Feed
	var feedResult FeedList

	db.Find(&groups)

	feedResult.Feeds = groups

	json.NewEncoder(w).Encode(feedResult)
}

func AddFeed(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	body := database.FeedBody{}

	err := decoder.Decode(&body)

	// Invalid body
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Missing params
	if body.Name == "" || body.Url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Everything OK, we store the feed and return 201
	_, err = database.CreateFeed(body, w)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)

}
func DeleteFeed(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["feedId"]

	id, err := strconv.ParseUint(idStr, 10, 32)

	fmt.Printf("Id : %s", idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.DeleteFeed(uint(id), w)

	if err != nil {
		return
	}

}

func ListFeedsInGroup(w http.ResponseWriter, r *http.Request) {

}

func AddFeedToGroup(w http.ResponseWriter, r *http.Request) {

}

func RemoveFeedFromGroup(w http.ResponseWriter, r *http.Request) {

}
