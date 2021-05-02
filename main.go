package main

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	id  int
	url string
}

type Group struct {
	gorm.Model
	id           int
	name         string
	feedsInGroup []Feed `gorm:many2many:group_feed`
}

func main() {

}

func handleRequests() {
}

func InitialMigration() {

}
