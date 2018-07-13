package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	"sync"
	"fmt"
)

var (
	once    sync.Once
	session *mgo.Session
)

func New() *mgo.Session {
	once.Do(func() {
		fmt.Println("Initializing MongoDB connection...")

		var err error

		session, err = mgo.Dial("localhost:27017")
		session.SetMode(mgo.Monotonic, true)
		if err != nil {
			log.Fatal(err)
		}
	})
	return session
}