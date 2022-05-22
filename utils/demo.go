package utils

import (
	"gopkg.in/mgo.v2"
	"log"
)

var context map[string]interface{}

func GetDemoContext() map[string]interface{} {
	return context
}

func SetDemoContext(name string, val interface{}) {
	context[name] = val
}

func init() {
	context = make(map[string]interface{})
	var err error

	// productIndex := mgo.Index{
	// 	Key:        []string{"code"},
	// 	Unique:     true,
	// 	Background: true,
	// 	Sparse:     true,
	// }

	userIndex := mgo.Index{
		Key:        []string{"username", "userid"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

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

	brandCol := session.DB(databaseName).C("demo_brands")
	err = brandCol.EnsureIndex(brandIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	modelCol := session.DB(databaseName).C("demo_models")
	err = modelCol.EnsureIndex(modelIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	dateCol := session.DB(databaseName).C("demo_dates")
	err = dateCol.EnsureIndex(dateIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	counterCol := session.DB(databaseName).C("demo_counter")
	err = counterCol.EnsureIndex(counterIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

}
