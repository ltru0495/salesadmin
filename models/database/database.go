package database

import (
	"admin/config"
	"admin/utils"

	"gopkg.in/mgo.v2/bson"
)

func insert(collection string, model interface{}) error {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(collection)

	err := C.Insert(&model)
	return err
}

func delete(collection string, fieldName string, fieldValue string) error {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(collection)

	err := C.Remove(bson.M{fieldName: fieldValue})
	return err
}
