package utils

import (
	"admin/config"
	"log"

	"gopkg.in/mgo.v2"
)

var msession *mgo.Session

var databaseName string

func init() {
	createSession()
	testPingandClose()
	addIndexes()
}

func createSession() {
	database := config.GetDatabaseConfig()
	session, err := mgo.Dial(database.Host)
	if err != nil {
		log.Println(err)
		return
	}
	msession = session
	msession.SetMode(mgo.Monotonic, true)

	databaseName = database.Database
}

func GetSession() *mgo.Session {
	if msession == nil {
		createSession()
	}
	return msession
}

func testPingandClose() {
	session := msession.Copy()
	defer session.Close()
	if err := session.Ping(); err != nil {
		log.Println(err)
	}
}

func CountCollection(collectionName string) int {
	session := msession.Copy()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	cant, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return cant
}

func CloseSession() {
	msession.Close()
}

func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"username", "userid"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	// productIndex := mgo.Index{
	// 	Key:        []string{"code"},
	// 	Unique:     true,
	// 	Background: true,
	// 	Sparse:     true,
	// }

	brandIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	modelIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	dateIndex := mgo.Index{
		Key:        []string{"date"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	counterIndex := mgo.Index{
		Key:        []string{"code"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	// Add indexes into MongoDB
	session := GetSession().Copy()
	defer session.Close()
	userCol := session.DB(databaseName).C("users")
	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	// prodCol := session.DB(databaseName).C("products")
	// err = prodCol.EnsureIndex(productIndex)
	// if err != nil {
	// 	log.Fatalf("[addIndexes]: %s\n", err)
	// }

	brandCol := session.DB(databaseName).C("brands")
	err = brandCol.EnsureIndex(brandIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	modelCol := session.DB(databaseName).C("models")
	err = modelCol.EnsureIndex(modelIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	dateCol := session.DB(databaseName).C("dates")
	err = dateCol.EnsureIndex(dateIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	counterCol := session.DB(databaseName).C("counter")
	err = counterCol.EnsureIndex(counterIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

}
