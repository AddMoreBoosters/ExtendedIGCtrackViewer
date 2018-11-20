package main 

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"fmt"
)

//	Initializes the database
func (db *TrackDB) Init () {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
}

//	Adds a track to the database
func (db *TrackDB) Add (track Track) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C("Placeholder").Insert(track)
	if err != nil {
		fmt.Printf("Error in Insert(): %v", err.Error())
	}
}

//	Gets a track from the database
func (db *TrackDB) Get (keyID string) (Track, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	track := Track{}
	err = session.DB(db.DatabaseName).C("Placeholder").Find(bson.M{"trackid": keyID}).One(&track)
	if err != nil {
		fmt.Printf("Error in Find().One(): %v", err.Error())
		return track, false
	}

	return track, true
}

//	Counts the tracks in the database
func (db *TrackDB) Count () int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	count, err := session.DB(db.DatabaseName).C("Placeholder").Count()
	if err != nil {
		fmt.Printf("Error in Count(): %v", err.Error())
		return -1
	}
	return count
}