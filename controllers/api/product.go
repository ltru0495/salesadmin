package api

import (
	"admin/models"
	"admin/models/database"
	"admin/utils"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"

	"log"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	code := params["code"]
	product, err := database.GetProductByCode(code)
	if err != nil {
		log.Println(err)
		return
	}

	j, err := json.Marshal(product)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func ProductChecked(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]
	err := database.SetProductChecked(code)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func GetLastProducts(w http.ResponseWriter, r *http.Request) {

	products, err := database.GetTodayProducts()
	if err != nil {
		log.Println(err)
		return
	}

	j, err := json.Marshal(products)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetBrands(w http.ResponseWriter, r *http.Request) {

	brands := database.GetAllBrands()

	j, err := json.Marshal(brands)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetModels(w http.ResponseWriter, r *http.Request) {

	models := database.GetAllModels()
	j, err := json.Marshal(models)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//  ********************* FOR CHARTS ************************
func GetBrandsForChart(w http.ResponseWriter, r *http.Request) {

	sales := database.GetSalesForCharts()

	brands := models.GetBrands(sales)
	j, err := json.Marshal(brands)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetSeriesForChart(w http.ResponseWriter, r *http.Request) {

	sales := database.GetSalesForCharts()

	series := models.GetSeries(sales)
	j, err := json.Marshal(series)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has ocurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

/* ============= FOR FILE =============== */
func ProductsBarcodeFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	var products []models.Product
	err = json.Unmarshal([]byte(r.PostFormValue("data")), &products)
	if err != nil {
		log.Println(err)
	}

	filename := "./public/codigos.pdf"

	var pdf *gofpdf.Fpdf
	pdf = models.BarcodesWithSizeFile(products)

	err = pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Println(err)
	}

	res := models.CreateDefaultResponse(w)
	res.Message = "/public/codigos.pdf"
	res.Send()
}

func ProductsBarcodeFile2(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	var products []models.Product
	err = json.Unmarshal([]byte(r.PostFormValue("data")), &products)
	if err != nil {
		log.Println(err)
	}

	filename := "./public/codigos.pdf"

	var pdf *gofpdf.Fpdf
	pdf = models.BarcodesFile(products)

	err = pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Println(err)
	}

	res := models.CreateDefaultResponse(w)
	res.Message = "/public/codigos.pdf"
	res.Send()
}
