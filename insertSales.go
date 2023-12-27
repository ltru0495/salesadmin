package main

import (
	"admin/models/database"
	"encoding/csv"
	"admin/models"
	"time"
	"os"
	"io"
	"log"
	"fmt"
	"strconv"
)


func main() {

//	var user models.User
//	user.Username = "admin2"
//	user.Password = "admin"
//	user.Name = "administrador"
//	user.Role = "admin"

//	database.InsertUser(&user)

	f, err := os.Open("vent.csv")

	    if err != nil {

        	log.Fatal(err)
	}

	r := csv.NewReader(f)
    for {

        record, err := r.Read()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatal(err)
        }
	dateString := record[2] + " " + record[3]
	timestamp, err := time.Parse("02/01/2006 15:04:05", dateString)
	if err!= nil {
		fmt.Println(err)
	}
	dateString = record[1]
	regDate, err := time.Parse("02/01/2006", dateString)
	if err != nil {
		fmt.Println(err)
	}
//	fmt.Println(record[0])
//	fmt.Println(date)

	size, err := strconv.Atoi(record[6])
	if err!= nil {
		fmt.Println(err)
	}

	priceB, err := strconv.ParseFloat(record[8], 32)
	if err != nil {
		fmt.Println(err)
	}
	priceS, err := strconv.ParseFloat(record[12], 32)
	if err != nil {
		fmt.Println(err)
	}


//	for value := range record {
//		 fmt.Printf("%s **", record[value])
//        }

	var p models.Sale

	if len(record) >= 15{
		p.PNote = record[14]
	}
	if len(record) == 16 {
		p.Comment = record[15]
	}

	p.Code = record[4]
	p.Category = ""
	p.Brand = record[5]
	p.Size = size
	p.Timestamp = timestamp
	p.Model = record[7]
	p.PriceBuy = priceB
	p.Price = priceS
	p.Location = record[9]
	p.Seller = record[10]
	p.RegDate = regDate
	p.Place = record[11]
	p.Payment_Method = record[13]

	fmt.Println(p.RegDate)
	fmt.Println(p.Timestamp)
	fmt.Println("**********************")
	err =	database.InsertSale2(&p)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(p)
	}

	}


}

