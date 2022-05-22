package database

import (
	"admin/config"
	"admin/models"
	"admin/utils"
	"time"
)

const (
	BRAND_COLLECTION = "brands"
	MODEL_COLLECTION = "models"
	DATE_COLLECTION  = "dates"
	NAME             = "name"
)

func InsertBrand(brand *models.Brand) error {
	err := insert(BRAND_COLLECTION, brand)
	return err
}

func InsertModel(model *models.Model) error {
	err := insert(MODEL_COLLECTION, model)
	return err
}

func InsertDate(date *models.Date) error {
	err := insert(DATE_COLLECTION, date)
	return err
}

func GetAllBrands() []models.Brand {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(BRAND_COLLECTION)

	var brands []models.Brand
	err := C.Find(nil).All(&brands)
	if err != nil {
		println(err)
	}
	return brands
}

func GetAllModels() []models.Model {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(MODEL_COLLECTION)

	var models []models.Model
	err := C.Find(nil).All(&models)
	if err != nil {
		println(err)
	}
	return models
}

func GetAllDates() []models.Date {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(DATE_COLLECTION)

	var dates []models.Date
	err := C.Find(nil).All(&dates)
	if err != nil {
		println(err)
	}

	var aux models.Date
	for k := 0; k < len(dates); k++ {
		for j := k + 1; j < len(dates); j++ {
			layout := "02-01-2006" // dd-MM-yyyy
			tk, err := time.Parse(layout, dates[k].Date)
			tj, err := time.Parse(layout, dates[j].Date)
			if err != nil {
				return dates
			}
			if tk.Before(tj) {
				aux = dates[k]
				dates[k] = dates[j]
				dates[j] = aux
			}
		}
	}
	return dates
}
