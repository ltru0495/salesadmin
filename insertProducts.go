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

	f, err := os.Open("inv.csv")

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
	dateString := record[0]
	date, err := time.Parse("02/01/2006", dateString)


	if err!= nil {
		fmt.Println(err)
	}
//	fmt.Println(record[0])
//	fmt.Println(date)

	size, err := strconv.Atoi(record[3])
	if err!= nil {
		fmt.Println(err)
	}

	priceS, err := strconv.ParseFloat(record[5], 32)
	if err != nil {
		fmt.Println(err)
	}
	priceB, err := strconv.ParseFloat(record[6], 32)
	if err != nil {
		fmt.Println(err)
	}


//	for value := range record {
//		 fmt.Printf("%s **", record[value])
//        }

	var p models.Product

	if len(record) == 9{
		p.Note = record[8]
	}

	p.Code = record[1]
	p.Brand = record[2]
	p.Size = size
	p.Price = priceB
	p.SPrice = priceS
	p.Location = record[7]
//	p.Note = record[8]
	p.RegDate = date
	p.Quantity = 1

	err = database.InsertProduct(&p)
	if err != nil {
		fmt.Println(err)
	}	else {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAa")
	}

//	fmt.Println(p.RegDate)
	}
}


