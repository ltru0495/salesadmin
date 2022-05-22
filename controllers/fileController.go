package controllers

import (
	"fmt"

	"github.com/gorilla/mux"

	"log"
	"net/http"
	"strings"
	"time"

	"io"
	"os"
	"strconv"

	"github.com/jung-kurt/gofpdf"

	"admin/models"
	"admin/models/database"
	"admin/utils"
)

func GetSalesFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}

	sales, err := database.GetSalesByDate(t)
	total, totalSales, sellers := models.GetTotalSales(sales)

	filename := "ventas_" + date + ".xlsx"
	file := models.GetReportFile(total, totalSales, sellers, sales, date, "all")

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
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)

	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetSalesPDFFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}

	sales, err := database.GetSalesByDate(t)
	total, totalSales, sellers := models.GetTotalSales(sales)

	filename := "./public/ventas_" + date + ".pdf"

	var pdf *gofpdf.Fpdf
	pdf = models.GetReportPDFFile(total, totalSales, sellers, sales, date, "all")

	err = pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Println(err)
	}

	Openfile, err := os.Open(filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		http.Error(w, "File not found.", 404)
		return
	}
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)

	// err = os.Remove(filename)
	// if err != nil {
	// 	log.Println(err)
	// }
	return
}

func GetSalesFileByLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}
	place := strings.ToUpper(params["place"])

	sales, err := database.GetSalesByDateAndLocation(t, place)
	total, totalSales, sellers := models.GetTotalSales(sales)

	filename := "ventas_" + date + ".xlsx"
	file := models.GetReportFile(total, totalSales, sellers, sales, date, place)

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
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)

	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetSalesPDFFileByLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}
	place := strings.ToUpper(params["place"])

	sales, err := database.GetSalesByDateAndLocation(t, place)
	total, totalSales, sellers := models.GetTotalSales(sales)

	filename := "ventas_" + date + ".pdf"
	var pdf *gofpdf.Fpdf
	pdf = models.GetReportPDFFile(total, totalSales, sellers, sales, date, "all")

	err = pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Println(err)
	}

	Openfile, err := os.Open(filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		http.Error(w, "File not found.", 404)
		return
	}
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)

	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetSalesFileForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	usr := utils.GetUser(r)
	place := usr.Name

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}

	sales, err := database.GetSalesByDateAndLocation(t, place)
	sellers := models.GetSellers(sales)

	filename := "ventas_" + date + ".xlsx"
	file := models.GetReportFileForUser(sellers, sales, date, place)

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
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)

	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetInventoryFile(w http.ResponseWriter, r *http.Request) {
	products := database.GetAllProducts()

	filename := "inventario.xlsx"
	file := models.InventoryFile(products)

	err := file.Save(filename)
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

func GetUncheckedProductsFile(w http.ResponseWriter, r *http.Request) {
	products, err := database.GetUncheckedProducts()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found.", 404)
		return
	}

	filename := "productos_faltantes.xlsx"

	file := models.GetInventoryFile(products)

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

func GetLastProductsFile(w http.ResponseWriter, r *http.Request) {

	t := time.Now()
	date := t.Format("02-01-2006")
	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}
	filename := "productos_" + date + ".xlsx"
	products, err := database.GetTodayProducts()
	if err != nil {
		log.Println(err)
		return
	}

	file := models.GetInventoryFile(products)

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
	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)

	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetLastProductsBarcodesFile(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	date := t.Format("02-01-2006")
	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}
	filename := "productos_codigos_" + date + ".pdf"
	products, err := database.GetTodayProducts()
	if err != nil {
		log.Println(err)
		return
	}
	params := mux.Vars(r)
	withSize := params["type"]

	var pdf *gofpdf.Fpdf
	pdf = models.LastBarcodesFile(products)
	if withSize == "s" {
		pdf = models.LastBarcodesFile(products)
	} else if withSize == "n" {
		pdf = models.BarcodesFile(products)
	} else if withSize == "g" {
		pdf = models.GroupedBarcodesFile(products)
	}

	err = pdf.OutputFileAndClose("./public/" + filename)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/public/"+filename, http.StatusSeeOther)

	return

}

func GetBarcodesFile(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	withSize := r.URL.Query().Get("size")

	code := r.URL.Query().Get("code")
	prod, _ := database.GetProductByCode(code)
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	for i := 0; i < n; i++ {
		products = append(products, prod)
	}

	filename := "codigos_" + prod.Code + ".pdf"
	var pdf *gofpdf.Fpdf
	if withSize == "s" {
		pdf = models.BarcodesWithSizeFile(products)
	} else {
		pdf = models.BarcodesFile(products)
	}
	err := pdf.OutputFileAndClose("./public/" + filename)
	if err != nil {
		log.Println(err)
	}
	res := models.CreateDefaultResponse(w)
	res.Message = "/public/" + filename
	res.Send()

	return
}
