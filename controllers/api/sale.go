package api

import (
	"admin/models"
	"admin/models/database"
	"admin/utils"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func GetSale(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	code := params["code"]
	sale, err := database.GetSaleByCode(code)
	if err != nil {
		log.Println(err)
		return
	}

	if sale.Price < 0 {
		log.Println(err)
		return
	}

	j, err := json.Marshal(sale)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func QuerySales(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	location := strings.ToUpper(params["location"])
	start, err := time.Parse("2006-01-02", params["start"])
	end, err := time.Parse("2006-01-02", params["end"])

	if err != nil {
		log.Println(err)
		return
	}

	sales, err := database.FilterSalesByDateAndLocation(start, end, location)
	if err != nil {
		log.Println(err)
		return
	}

	j, err := json.Marshal(sales)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
func GetSalesFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	location := strings.ToUpper(params["location"])
	start, err := time.Parse("2006-01-02", params["start"])
	end, err := time.Parse("2006-01-02", params["end"])

	if err != nil {
		log.Println(err)
		return
	}

	sales, err := database.FilterSalesByDateAndLocation(start, end, location)
	if err != nil {
		log.Println(err)
		return
	}
	
	filename := "REPORTE DE VENTAS.xlsx"
	file := models.SaleFile(sales)

	err = file.Save(filename)
	if err != nil {
		fmt.Printf(err.Error())
	}

	Openfile, err := os.Open(filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		http.Error(w, "File not found.", 404)
		return
	}
	models.SendFile(w, Openfile, filename)

	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	return
}
