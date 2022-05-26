package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Sale struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code     string        `json:"code" bson:"code"`
	Brand    string        `json:"brand,omitempty" bson:"brand,omitempty"`
	Serie    string        `json:"serie,omitempty" bson:"serie,omitempty"`
	Size     int           `json:"size,omitempty" bson:"size,omitempty"`
	Model    string        `json:"model,omitempty" bson:"model,omitempty"`
	Location string        `json:"location,omitempty" bson:"location,omitempty"`

	PriceBuy float64 `json:"pricebuy,omitempty" bson:"pricebuy,omitempty"`
	Earning  float64 `json:"earning,omitempty" bson:"earning,omitempty"`

	Seller    string    `json:"seller" bson:"seller"`
	Place     string    `json:"place" bson:"place"` //Lugar de venta
	Price     float64   `json:"price" bson:"price"`
	
	Comment   string    `json:"comment" bson:"comment"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	RegDate   time.Time `json:"regdate" bson:"regdate"`

	Refunded bool `json:"refunded,omitempty" bson:"refunded,omitempty"`

	Payment_Method string `json:"payment_method" bson:"payment_method"`
}

func (s *Sale) ToMap() map[string]string {
	m := make(map[string]string)
	j, _ := json.Marshal(s)
	json.Unmarshal(j, &m)
	m["pricebuy"] = fmt.Sprintf("%.02f", s.PriceBuy)
	m["price"] = fmt.Sprintf("%.02f", s.Price)
	m["size"] = fmt.Sprintf("%d", s.Size)

	m["regdate"] = spanishDate(s.RegDate.Format("02/Jan/2006"))
	if m["regdate"] == "01/Ene/0001" {
		m["regdate"] = ""
	}
	m["timestamp"] = spanishDate(s.Timestamp.Format("02/Jan/2006"))
	if m["timestamp"] == "01/Ene/0001" {
		m["timestamp"] = ""
	}
	m["time"] = spanishDate(s.Timestamp.Format("15:04:05"))
	return m
}

type Date struct {
	Date string `json:"date" bson:"date"`
}

type Seller struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func (s Sale) GetFullTimestamp() string {
	date := s.Timestamp.Format("02-01-2006 15:04")
	return date
}

func (s Sale) GetTimestamp() string {
	date := s.Timestamp.Format("15:04")
	return date
}

func (s Sale) IsRefund() bool {
	if s.Price < 0 {
		return true
	}
	return false
}

func (sale Sale) GetProduct() (product Product) {
	product = Product{
		Code:     sale.Code,
		Brand:    sale.Brand,
		Serie:    sale.Serie,
		Size:     sale.Size,
		Model:    sale.Model,
		Location: sale.Location,
		Price:    sale.PriceBuy,
		Quantity: 1,
	}
	product.SetRegistrationDate()
	product.SetModificationDate()

	return
}

func GetTotal(sales []Sale) float64 {
	total := 0.0
	for _, sale := range sales {
		total = total + sale.Price
	}
	return total
}

func GetEarning(sales []Sale) float64 {
	total := 0.0
	for _, sale := range sales {
		total = total + sale.Earning
	}
	return total
}

func GetTotalSales(sales []Sale) (float64, float64, []Seller) {
	var s map[string]int
	s = make(map[string]int)

	sellers := GetSellers(sales)

	total := 0.0
	totalSales := 0.0
	for _, sale := range sales {
		total = total + sale.Earning
		totalSales = totalSales + sale.Price
	}

	for k, v := range s {
		if k != "" {
			sellers = append(sellers, Seller{k, v})
		}
	}

	return total, totalSales, sellers
}
func GetTotalSalesByPM(sales []Sale, paymentMethod string) float64 {
	total := 0.0
	for _, sale := range sales {
		if sale.Payment_Method == paymentMethod {
			total = total + sale.Price
		}
	}
	return total
}

func GetSellers(sales []Sale) []Seller {
	var s map[string]int
	s = make(map[string]int)

	var sellers []Seller

	for _, sale := range sales {
		seller := strings.TrimSpace(sale.Seller)
		if _, ok := s[seller]; ok {
			if strings.ToUpper(sale.Comment) == "DEVOLUCIÓN" || strings.ToUpper(sale.Comment) == "DEVOLUCION" {
				s[seller] = s[seller] - 1
			} else if strings.ToUpper(sale.Comment) == "CAMBIO" {
				s[seller] = s[seller] + 0
			} else {
				s[seller] = s[seller] + 1
			}
		} else {
			if strings.ToUpper(sale.Comment) == "DEVOLUCIÓN" || strings.ToUpper(sale.Comment) == "DEVOLUCION" {
				s[seller] = s[seller] - 1
			} else if strings.ToUpper(sale.Comment) == "CAMBIO" {
				s[seller] = s[seller] + 0
			} else {
				s[seller] = s[seller] + 1
			}
		}
	}

	for k, v := range s {
		if k != "" {
			sellers = append(sellers, Seller{k, v})
		}
	}

	return sellers
}

func GetSalesAndRefunds(s []Sale) ([]Sale, []Sale) {
	var sales []Sale
	var refunds []Sale

	for _, sale := range s {
		if strings.ToUpper(sale.Comment) == "DEVOLUCIÓN" || strings.ToUpper(sale.Comment) == "CAMBIO" {
			refunds = append(refunds, sale)
		} else {
			sales = append(sales, sale)
		}
	}

	return sales, refunds
}

func GetBrands(sales []Sale) []Brand {
	var s map[string]int
	s = make(map[string]int)

	var brands []Brand

	for _, sale := range sales {
		if _, ok := s[sale.Brand]; ok {
			s[sale.Brand] = s[sale.Brand] + 1
		} else {
			s[sale.Brand] = 1
		}
	}

	for k, v := range s {
		if k != "" {
			brands = append(brands, Brand{k, v})
		}
	}

	var aux Brand
	for i := 0; i < len(brands); i++ {
		for k := i + 1; k < len(brands); k++ {
			if brands[i].Amount <= brands[k].Amount {
				aux = brands[i]
				brands[i] = brands[k]
				brands[k] = aux
			}
		}
	}

	if len(brands) > 30 {
		brands = brands[0:30]
	}
	return brands
}

// func GetSeries(sales []Sale) []Serie {
// 	// series := []string{"21-26", "27-32", "33-38", "39-44"}
// 	var s map[string]int
// 	s = make(map[string]int)
// 	s["21-26"] = 0
// 	s["27-32"] = 0
// 	s["33-38"] = 0
// 	s["39-44"] = 0

// 	var series []Serie

// 	for _, sale := range sales {
// 		if sale.Size >= 21 && sale.Size <= 26 {
// 			s["21-26"] += 1
// 		} else if sale.Size >= 27 && sale.Size <= 32 {
// 			s["27-32"] += 1
// 		} else if sale.Size >= 33 && sale.Size <= 38 {
// 			s["33-38"] += 1
// 		} else if sale.Size >= 39 && sale.Size <= 44 {
// 			s["39-44"] += 1
// 		}
// 	}

// 	for k, v := range s {
// 		if k != "" {

// 			series = append(series, Serie{k, v})
// 		}
// 	}

// 	return series
// }
func GetSeries(sales []Sale) []Serie {
	// series := []string{"21-26", "27-32", "33-38", "39-44"}
	var s map[string]int
	s = make(map[string]int)
	for i := 22; i <= 44; i++ {
		s[strconv.Itoa(i)] = 0
	}

	var series []Serie

	for _, sale := range sales {
		s[strconv.Itoa(sale.Size)] += 1
	}

	for k, v := range s {
		if k != "" {
			series = append(series, Serie{k, v})
		}
	}

	return series
}

func SortSales(sales []Sale) []Sale {
	var aux Sale
	for k := 0; k < len(sales); k++ {
		for j := k + 1; j < len(sales); j++ {
			tk := sales[k].Timestamp
			tj := sales[j].Timestamp
			if tk.After(tj) {
				aux = sales[k]
				sales[k] = sales[j]
				sales[j] = aux
			}
		}
	}
	return sales
}
