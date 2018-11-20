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

//	Adds a track to the database, and returns its new id
func (db *TrackDB) Add (track Track) string {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C("Placeholder").Insert(track)
	if err != nil {
		fmt.Printf("Error in Insert(): %v", err.Error())
	}

	err = session.DB(db.DatabaseName).C("Placeholder").Find(bson.M{"track_src_url": track.TrackSourceURL}).One(&track)
	if err != nil {
		fmt.Printf("Error in Find().One(): %v", err.Error())
	}
	return track._id
}

//	Gets a track from the database
func (db *TrackDB) Get (keyID string) (Track, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	track := Track{}
	err = session.DB(db.DatabaseName).C("Placeholder").Find(bson.M{"_id": keyID}).One(&track)
	if err != nil {
		fmt.Printf("Error in Find().One(): %v", err.Error())
		return track, false
	}

	return track, true
}

//	Gets all indexes from the database
func (db *TrackDB) GetIndexes () []string {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var docs []Track
	err = session.DB(db.DatabaseName).C("Placeholder").Find().All(&docs)
	if err != nil {
		fmt.Printf("Error in Find().All(): %v", err.Error())
		return nil
	}

	var indexes []string
	for _, doc := range docs {
		indexes = append(indexes, doc._id)
	}
	return indexes
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