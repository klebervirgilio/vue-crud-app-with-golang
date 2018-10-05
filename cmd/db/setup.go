package main

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

func main() {
	var err error
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://mongo_user:mongo_secret@0.0.0.0:27017"
	}
	session, err := mgo.Dial(mongoURL)
	defer session.Close()

	err = session.DB("").AddUser("mongo_user", "mongo_secret", false)
	// err = session.DB("kudos_test").AddUser("mongo_user", "mongo_secret", false)

	info := &mgo.CollectionInfo{}
	err = session.DB("").C("kudos").Create(info)
	// err = session.DB("kudos_test").C("kudos").Create(info)

	if err != nil {
		log.Fatal(err)
	}
}
