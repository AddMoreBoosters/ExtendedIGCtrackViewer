package main 

import (
	"testing"
	"github.com/globalsign/mgo"
)

func setupDB() *TrackDB {
	db := TrackDB {
		"mongodb://localhost",
		"testTrackDB",
		"tracks",
	}

	session, err := mgo.Dial(db.DatabaseURL)
}

func tearDownDB(db *TrackDB) {

}

func TestTrackMongoDB_Add (t *testing.T) {

}