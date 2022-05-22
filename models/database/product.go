package database

import (
	"admin/config"
	"admin/models"
	"admin/utils"

	"gopkg.in/mgo.v2/bson"

	// "errors"
	"time"
	// "log"
)

const (
	PRODUCT_COLLECTION = "products"
	COUNTER_COLLECTION = "counter"
	CODE               = "code"
	REGDATE            = "regdate"

	COUNTERCODE = "MyCounter"
)

func InsertProduct(product *models.Product) error {
	product.SetRegistrationDate()
	product.SetUnchecked()
	err := insert(PRODUCT_COLLECTION, product)
	return err
}

func InsertProductRefund(product *models.Product) error {
	product.SetUnchecked()
	err := insert(PRODUCT_COLLECTION, product)
	return err
}

func InsertCounter(counter *models.Counter) error {
	counter.Timestamp = time.Now()
	err := insert(COUNTER_COLLECTION, counter)
	return err
}

func UpdateCounter() error {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(COUNTER_COLLECTION)

	counter := models.Counter{}
	err := C.Find(bson.M{CODE: COUNTERCODE}).One(&counter)
	if err != nil {
		return err
	}
	counter.Counter = counter.Counter + 1
	err = C.Update(bson.M{CODE: COUNTERCODE}, counter)
	return err
}

func GetCounter() (models.Counter, error) {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(COUNTER_COLLECTION)

	counter := models.Counter{}
	err := C.Find(bson.M{CODE: COUNTERCODE}).One(&counter)

	return counter, err

}

func GetProductByCode(code string) (models.Product, error) {
	product := models.Product{}
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	err := C.Find(bson.M{CODE: code}).One(&product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetTodayProducts() ([]models.Product, error) {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	t := time.Now()
	year, month, day := t.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	next := date.AddDate(0, 0, 1)
	dateUTC := date.UTC()
	nextUTC := next.UTC()
	_, offset := nextUTC.Local().Zone()

	dateUTC = dateUTC.Add(-time.Duration(offset) * time.Second)
	nextUTC = nextUTC.Add(-time.Duration(offset) * time.Second)
	var products []models.Product
	err := C.Find(bson.M{REGDATE: bson.M{
		"$gt": dateUTC,
		"$lt": nextUTC,
	}}).All(&products)
	if err != nil {
		return products, err
	}
	return products, nil
}

func UpdateProductByCode(code string, product *models.Product) error {
	productToUpdate := models.Product{}

	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	err := C.Find(bson.M{CODE: code}).One(&productToUpdate)
	if err != nil {
		return err
	}
	product.Quantity = productToUpdate.Quantity
	product.RegDate = productToUpdate.RegDate
	product.SetModificationDate()

	err = C.Update(bson.M{CODE: code}, product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductByCode(code string) error {
	err := delete(PRODUCT_COLLECTION, CODE, code)
	return err
}

func GetAllProducts() []models.Product {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	var products []models.Product
	err := C.Find(nil).All(&products)
	if err != nil {
		println(err)
	}
	return products
}

func UpdateProductQuantity(code string) error {
	productToUpdate := models.Product{}

	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	err := C.Find(bson.M{CODE: code}).One(&productToUpdate)
	if err != nil {
		return err
	}

	productToUpdate.Quantity = productToUpdate.Quantity - 1

	if productToUpdate.Quantity <= 0 {
		DeleteProductByCode(productToUpdate.Code)
		return nil
	}

	err = C.Update(bson.M{CODE: code}, productToUpdate)
	if err != nil {
		return err
	}
	return nil
}

func SetProductChecked(code string) error {
	productToUpdate := models.Product{}

	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	err := C.Find(bson.M{CODE: code}).One(&productToUpdate)
	if err != nil {
		return err
	}
	productToUpdate.Check = true

	err = C.Update(bson.M{CODE: code}, productToUpdate)
	if err != nil {
		return err
	}
	return nil
}

func GetUncheckedProducts() ([]models.Product, error) {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	var products []models.Product
	err := C.Find(bson.M{"check": false}).All(&products)

	if err != nil {
		return products, err
	}
	return products, nil
}

func SetAllProductsUnchecked() error {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(PRODUCT_COLLECTION)

	products := GetAllProducts()
	var err error
	for _, product := range products {
		product.Check = false
		err = C.Update(bson.M{CODE: product.Code}, product)
	}
	return err
}
