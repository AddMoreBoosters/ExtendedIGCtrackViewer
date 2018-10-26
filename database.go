package main 

import (
	"github.com/globalsign/mgo"
)

type TrackDB struct {
	DatabaseURL string
	DatabaseName string
}

func (db *TrackDB) Init () {

}

func (db *TrackDB) Add (track Track) {
	
}

func (db *TrackDB) Get (keyID string) {
	
}

func (db *TrackDB) Count () int {
	
}