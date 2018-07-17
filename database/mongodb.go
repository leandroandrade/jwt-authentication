package database

import (
	"gopkg.in/mgo.v2"
	"log"
	"sync"
)

var (
	once    sync.Once
	session *mgo.Session
)

type MongoDatabase struct {
}

func New() *MongoDatabase {
	once.Do(func() {
		log.Println("Initializing MongoDB connection...")

		var err error

		session, err = mgo.Dial("localhost:27017")
		session.SetMode(mgo.Monotonic, true)
		if err != nil {
			log.Fatal(err)
		}
	})
	return &MongoDatabase{}
}

func (MongoDatabase) Get() *mgo.Session {
	return session.Copy()
}
