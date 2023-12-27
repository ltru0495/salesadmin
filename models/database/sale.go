package database

import (
	"admin/config"
	"admin/models"
	"admin/utils"
	"time"

	"gopkg.in/mgo.v2/bson"

	"errors"
	"fmt"
)

const (
	SALE_COLLECTION = "sales"
	SALE_CODE       = "code"
	TIMESTAMP       = "timestamp"
	SELLER          = "seller"
	PLACE           = "place"
	BRAND           = "brand"
	ID              = "_id"
)

func InsertSale(sale *models.Sale) error {
	sale.Id = bson.NewObjectId()
	sale.Timestamp = time.Now()
	prod, err := GetProductByCode(sale.Code)
	if err != nil {
		return errors.New("Error al registrar venta: producto no encontrado")
	}

	sale.Brand = prod.Brand
	sale.Model = prod.Model
	sale.Size = prod.Size
	sale.Category = prod.Category
	sale.Refunded = false

	sale.PriceBuy = prod.Price
	sale.Serie = prod.Serie
	sale.Earning = sale.Price - prod.Price

	sale.RegDate = prod.RegDate
	err = insert(SALE_COLLECTION, sale)
	return err
}

func DeleteSaleById(id string) error {
	sale, err := GetSaleById(id)
	if err != nil {
		return err
	}
	product := sale.GetProduct()

	if sale.Price > 0 {
		err = InsertProduct(&product)
		if err != nil {
			fmt.Println(err)
			return errors.New("Error al recuperar el producto")
		}
	}

	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)
	err = C.Remove(bson.M{ID: bson.ObjectIdHex(id)})

	return err
}

func RefundSaleById(id string, action string) error {
	sale, err := GetSaleById(id)
	if err != nil {
		return err
	}

	product := sale.GetProduct()
	product.RegDate = time.Now().AddDate(0, 0, -3)

	err = InsertProductRefund(&product)

	if err != nil {
		return errors.New("Error al recuperar el producto")
	}
	sale.Id = bson.NewObjectId()
	sale.Timestamp = time.Now()
	sale.Comment = action
	sale.Price = -sale.Price
	sale.Earning = 0.0
	sale.Refunded = true

	err = insert(SALE_COLLECTION, sale)

	// session := utils.GetSession().Copy()
	// defer session.Close()
	// conf := config.GetDatabaseConfig()
	// dbname := conf.Database
	// C := session.DB(dbname).C(SALE_COLLECTION)
	// err = C.Remove(bson.M{ID: bson.ObjectIdHex(id)})

	// if err != nil {
	// 	return errors.New("Error al eliminar la venta")
	// }

	return err
}

func GetSaleById(id string) (models.Sale, error) {
	sale := models.Sale{}
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	err := C.Find(bson.M{ID: bson.ObjectIdHex(id)}).One(&sale)
	if err != nil {
		return models.Sale{}, err
	}
	return sale, nil
}

func GetSaleByCode(code string) (models.Sale, error) {
	sale := models.Sale{}
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)
	err := C.Find(bson.M{SALE_CODE: code}).One(&sale)
	if err != nil {
		return models.Sale{}, err
	}
	return sale, nil
}

func GetAllSales() []models.Sale {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	var sales []models.Sale
	err := C.Find(nil).All(&sales)
	if err != nil {
		println(err)
	}
	return sales
}

func GetSalesForCharts() []models.Sale {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	date := time.Date(2019, 5, 17, 0, 0, 0, 0, time.UTC)
	_, offset := date.Local().Zone()
	dateUTC := date.Add(-time.Duration(offset) * time.Second)

	var sales []models.Sale
	err := C.Find(bson.M{TIMESTAMP: bson.M{
		"$gt": dateUTC,
	}}).Sort("timestamp").All(&sales)
	if err != nil {
		return sales
	}

	return sales
}

func GetSalesByDateAndLocation(date time.Time, place string) ([]models.Sale, error) {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	next := date.AddDate(0, 0, 1)
	dateUTC := date.UTC()
	nextUTC := next.UTC()

	_, offset := nextUTC.Local().Zone()

	dateUTC = dateUTC.Add(-time.Duration(offset) * time.Second)
	nextUTC = nextUTC.Add(-time.Duration(offset) * time.Second)

	var sales []models.Sale
	err := C.Find(bson.M{TIMESTAMP: bson.M{
		"$gt": dateUTC,
		"$lt": nextUTC,
	}, PLACE: place}).Sort("timestamp").All(&sales)
	if err != nil {
		return sales, err
	}
	return sales, nil
}

func FilterSalesByDateAndLocation(start, end time.Time, place string) ([]models.Sale, error) {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	startUTC := start.UTC()
	endUTC := end.UTC()

	_, offset := endUTC.Local().Zone()

	startUTC = startUTC.Add(-time.Duration(offset) * time.Second)
	endUTC = endUTC.Add(-time.Duration(offset) * time.Second)

	var sales []models.Sale

	if place == "ALL" {
		err := C.Find(bson.M{TIMESTAMP: bson.M{
			"$gt": startUTC,
			"$lt": endUTC,
		}}).Sort("timestamp").All(&sales)
		if err != nil {
			return sales, err
		}
		return sales, nil
	}

	err := C.Find(bson.M{TIMESTAMP: bson.M{
		"$gt": startUTC,
		"$lt": endUTC,
	}, PLACE: place}).Sort("timestamp").All(&sales)
	if err != nil {
		return sales, err
	}
	return sales, nil
}

func GetSalesByDate(date time.Time) ([]models.Sale, error) {
	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	next := date.AddDate(0, 0, 1)
	dateUTC := date.UTC()
	nextUTC := next.UTC()

	_, offset := nextUTC.Local().Zone()

	dateUTC = dateUTC.Add(-time.Duration(offset) * time.Second)
	nextUTC = nextUTC.Add(-time.Duration(offset) * time.Second)

	var sales []models.Sale

	err := C.Find(bson.M{TIMESTAMP: bson.M{
		"$gt": dateUTC,
		"$lt": nextUTC,
	}}).Sort("timestamp").All(&sales)
	if err != nil {
		return sales, err
	}

	return sales, nil
}

func GetSalesOfTheMonth() ([]models.Sale, error) {

	y, m, d := time.Now().Date()
	firstDay := time.Date(y, m-1, d, 0, 0, 0, 0, time.UTC)
	lastDay := time.Date(y, m, d, 0, 0, 0, -1, time.UTC)

	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	dateUTC := firstDay.UTC()
	nextUTC := lastDay.UTC()

	var allsales []models.Sale
	var sales []models.Sale

	err := C.Find(bson.M{TIMESTAMP: bson.M{
		"$gt": dateUTC,
		"$lt": nextUTC,
	}}).Sort("timestamp").All(&allsales)

	for _, sale := range allsales {

		if sale.Price > 0 && !sale.Refunded {
			sales = append(sales, sale)
		}
	}

	if err != nil {
		return sales, err
	}

	return sales, nil
}

func UpdateSaleById(id string, sale *models.Sale) error {
	saleToUpdate := models.Sale{}

	session := utils.GetSession().Copy()
	defer session.Close()
	conf := config.GetDatabaseConfig()
	dbname := conf.Database
	C := session.DB(dbname).C(SALE_COLLECTION)

	err := C.Find(bson.M{ID: bson.ObjectIdHex(id)}).One(&saleToUpdate)
	if err != nil {
		return err
	}
	saleToUpdate.Price = sale.Price
	saleToUpdate.Earning = saleToUpdate.Price - saleToUpdate.PriceBuy
	saleToUpdate.Seller = sale.Seller
	saleToUpdate.Comment = sale.Comment
	saleToUpdate.Payment_Method = sale.Payment_Method
	err = C.Update(bson.M{ID: bson.ObjectIdHex(id)}, saleToUpdate)
	if err != nil {
		return err
	}
	return nil
}
