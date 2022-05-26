package models

import (
	"encoding/json"
	"fmt"
	"strings"

	// "gopkg.in/mgo.v2/bson"
	"time"
)

type Product struct {

	// Obligatorios
	Code     string  `json:"code" bson:"code"`
	Brand    string  `json:"brand" json:"brand"`
	Serie    string  `json:"serie" bson:"serie" `
	Size     int     `json:"size" bson:"size"`
	Model    string  `json:"model" bson:"model"`
	Price    float64 `json:"price" bson:"price"`
	Quantity int     `json:"quantity" bson:"quantity"`
	Location string  `json:"location" bson:"location"`
	PFC      string  `json:"pfc" bson:"pfc"` //Factory Product Code
	//

	Check bool `json:"check" bson:"check"`

	Note    string    `json:"note" bson:"note"`
	RegDate time.Time `json:"regdate" bson:"regdate"`
	ModDate time.Time `json:"moddate" bsoon:"moddate"`
}

func (p *Product) String() string {
	return fmt.Sprintf("Producto : %s - %s\nPrecio: %f\n", p.Code, p.Serie, p.Price)
}

func (p *Product) SetRegistrationDate() {
	p.RegDate = time.Now()
}

func (p *Product) SetModificationDate() {
	p.ModDate = time.Now()
}

func (p *Product) SetUnchecked() {
	p.Check = false
}

func spanishDate(date string) string {
	date = strings.Replace(date, "Jan", "Ene", 1)
	date = strings.Replace(date, "Apr", "Abr", 1)
	date = strings.Replace(date, "Aug", "Ago", 1)
	date = strings.Replace(date, "Dec", "Dic", 1)
	return date
}

func (p *Product) ToMap() map[string]string {
	m := make(map[string]string)
	j, _ := json.Marshal(p)
	json.Unmarshal(j, &m)
	m["price"] = fmt.Sprintf("%.02f", p.Price)
	m["size"] = fmt.Sprintf("%d", p.Size)
	m["regdate"] = spanishDate(p.RegDate.Format("02/01/2006"))
	return m
}

type Brand struct {
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount,omitempty" bson:"amount,omitempty"`
}

type Model struct {
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount" bson:"amount"`
}

type Serie struct {
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount" bson:"amount"`
}

type Counter struct {
	Code      string    `json:"code" bson:"code"`
	Timestamp time.Time `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Counter   int       `json:"counter" bson:"counter"`
}
