package controllers

import (
	"admin/models"
	"admin/models/database"
	"admin/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func NewSale(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			utils.SetContext(r, "Error", err)
			http.Redirect(w, r, "/venta/nueva", 301)
			return
		} else {
			sale := new(models.Sale)

			// GET from form
			decoder := schema.NewDecoder()
			err = decoder.Decode(sale, r.PostForm)

			if err != nil {
				log.Println(err)
				utils.SetContext(r, "Error", err)
				http.Redirect(w, r, "/venta/nueva", 301)
				return

			} else {
				//Save
				t := time.Now()
				date := new(models.Date)
				date.Date = t.Format("02-01-2006")
				err = database.InsertDate(date)
				if err != nil {
					log.Println(err)
					utils.SetContext(r, "Error", err)
					http.Redirect(w, r, "/venta/nueva", 301)
				}
				_, err := database.GetProductByCode(sale.Code)
				if err == nil {
					err = database.InsertSale(sale)
					if err != nil {
						log.Println(err)
						utils.SetContext(r, "Error", err)
						http.Redirect(w, r, "/venta/nueva", 301)
					} else {

						err = database.UpdateProductQuantity(sale.Code)
						if err != nil {
							log.Println(err)
							utils.SetContext(r, "Error", err)
							return
						} else {
							utils.SetContext(r, "Success", "Se registro venta correctamente")
							return
						}
					}
				}

			}
		}
		return
	}

	t := time.Now()
	date := t.Format("02-01-2006")

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}
	context := utils.GetFullContext(r)
	usr := utils.GetUser(r)

	sales, err := database.GetSalesByDateAndLocation(t, usr.Name)
	if err != nil {
		log.Println(err)
	}
	sellers := models.GetSellers(sales)
	sales, refunds := models.GetSalesAndRefunds(sales)
	context["Sales"] = sales
	context["Title"] = "Ventas del Día"
	context["Refunds"] = refunds
	context["Sellers"] = sellers
	context["IsIndex"] = true

	utils.RenderTemplate(w, "sale_new", context)
	utils.SetContext(r, "Error", nil)
	utils.SetContext(r, "Success", nil)
}

func Sales(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}

	context := utils.GetFullContext(r)

	usr := utils.GetUser(r)
	if usr.Role == "admin" {
		sales, err := database.GetSalesByDate(t)

		sellers := models.GetSellers(sales)
		if err != nil {
			log.Println(err)
		}

		sales, refunds := models.GetSalesAndRefunds(sales)
		if err != nil {
			log.Println(err)
		}

		salesC, err := database.GetSalesByDateAndLocation(t, "CAJAMARCA")
		salesC, refundsC := models.GetSalesAndRefunds(salesC)
		totalSC := models.GetTotal(salesC)
		totalRC := models.GetTotal(refundsC)

		salesL, err := database.GetSalesByDateAndLocation(t, "LORETO")
		salesL, refundsL := models.GetSalesAndRefunds(salesL)
		totalSL := models.GetTotal(salesL)
		totalRL := models.GetTotal(refundsL)

		// <td>{{ .TotalPEL }}</td>
		// <td>{{ .TotalEL }}</td>
		// <td>{{ .UtilidadL }}</td>

		context["Title"] = "Ventas del día"

		context["TotalSalesC"] = totalSC
		context["TotalRefundsC"] = totalRC
		context["TotalEarningC"] = models.GetEarning(salesC)
		context["TotalSalesAndRefundsC"] = totalSC + totalRC
		context["SalesC"] = salesC

		context["TotalPEC"] = models.GetTotalSalesByPM(salesC, "electronico")
		context["TotalEC"] = models.GetTotalSalesByPM(salesC, "efectivo")
		context["UtilidadC"] = models.GetEarning(salesC)

		context["TotalSalesL"] = totalSL
		context["TotalRefundsL"] = totalRL
		context["TotalEarningL"] = models.GetEarning(salesL)
		context["TotalSalesAndRefundsL"] = totalSL + totalRL
		context["SalesL"] = salesL

		context["TotalPEL"] = models.GetTotalSalesByPM(salesL, "electronico")
		context["TotalEL"] = models.GetTotalSalesByPM(salesL, "efectivo")
		context["UtilidadL"] = models.GetEarning(salesL)
		// totalS := models.GetTotal(sales)
		// totalR := models.GetTotal(refunds)
		// context["Title"] = "Ventas del día"
		// context["TotalSales"] = totalS
		// context["TotalRefunds"] = totalR
		// context["TotalEarning"] = models.GetEarning(sales)
		// context["TotalSalesAndRefunds"] = totalS + totalR
		context["Sellers"] = sellers
		// context["Sales"] = sales
		context["Refunds"] = refunds
		// sales, refunds := models.GetSalesAndRefunds(sales)

		wd := fmt.Sprintf("%v", t.Weekday())
		context["Title"] = "Ventas de " + GetWeekDay(wd) + " (" + date + ")"
		utils.RenderTemplate(w, "index_admin", context)
		utils.SetContext(r, "Sellers", nil)
		utils.SetContext(r, "Sales", nil)
		return
	} else {
		sales, err := database.GetSalesByDateAndLocation(t, usr.Name)

		sellers := models.GetSellers(sales)
		if err != nil {
			log.Println(err)
		}
		sales, refunds := models.GetSalesAndRefunds(sales)

		context["Sales"] = sales
		context["Sellers"] = sellers
		context["Refunds"] = refunds
	}
	wd := fmt.Sprintf("%v", t.Weekday())
	context["Title"] = "Ventas de " + GetWeekDay(wd) + " (" + date + ")"
	utils.RenderTemplate(w, "sales", context)
	utils.SetContext(r, "Sellers", nil)
	utils.SetContext(r, "Sales", nil)
}

func GetWeekDay(wd string) string {
	switch wd {
	case "Saturday":
		return "Sabado"
	case "Sunday":
		return "Domingo"
	case "Monday":
		return "Lunes"
	case "Tuesday":
		return "Martes"
	case "Wednesday":
		return "Miércoles"
	case "Thursday":
		return "Jueves"
	case "Friday":
		return "Viernes"
	default:
		return ""
	}

}

func SalesRefund(w http.ResponseWriter, r *http.Request) {
	context := utils.GetFullContext(r)

	// usr := utils.GetUser(r)
	// if usr.Role == "usr" {
	utils.RenderTemplate(w, "sale_refund_user", context)
	return
	// }

	// sales, err := database.GetSalesOfTheMonth()
	// sales = models.SortSales(sales)
	// if err != nil {
	// 	log.Println(err)
	// }
	// context["Sales"] = sales
	// utils.RenderTemplate(w, "sale_refund_user", context)

}

func SalesByLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	layout := "02-01-2006" // dd-MM-yyyy
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
		return
	}

	place := strings.ToUpper(params["place"])

	context := utils.GetFullContext(r)

	sales, err := database.GetSalesByDateAndLocation(t, place)
	total, totalSales, sellers := models.GetTotalSales(sales)

	sales, refunds := models.GetSalesAndRefunds(sales)
	if err != nil {
		log.Println(err)
	}

	context["Sellers"] = sellers
	context["Total"] = total
	context["TotalSales"] = totalSales
	context["Sales"] = sales
	context["Refunds"] = refunds

	context["Title"] = "Ventas de " + date + "               " + strings.Title(place)

	utils.RenderTemplate(w, "sales", context)
}

func DeleteSale(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		utils.SetContext(r, "Error", err)
	} else {
		password := r.PostFormValue("password")

		if err = database.CompareAdminPassword(password); err != nil {
			log.Println(err)
			utils.SetContext(r, "Error", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		err = database.DeleteSaleById(id)
		if err != nil {
			utils.SetContext(r, "Error", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		utils.SetContext(r, "Success", "Se eliminó la venta correctamente")
		w.WriteHeader(http.StatusOK)
	}
}

func RefundSale(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	actionParam := params["action"]
	var action string
	if actionParam == "devolver" {
		action = "DEVOLUCIÓN"
	}
	if actionParam == "cambiar" {
		action = "CAMBIO"
	}
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		utils.SetContext(r, "Error", err)
	} else {
		err = database.RefundSaleById(id, action)
		if err != nil {
			utils.SetContext(r, "Error", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		utils.SetContext(r, "Success", "Se registró la devolución la venta correctamente")
		w.WriteHeader(http.StatusOK)
	}
}

func ViewSale(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	sale, err := database.GetSaleById(id)
	if err != nil {
		log.Println(err)
		return
	}
	context := utils.GetFullContext(r)
	context["Sale"] = sale

	utils.RenderTemplate(w, "sale_view", context)
	utils.SetContext(r, "Success", nil)
	utils.SetContext(r, "Error", nil)
}

func UpdateSale(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		utils.SetContext(r, "Error", err)
	} else {
		sale := new(models.Sale)
		decoder := schema.NewDecoder()
		err = decoder.Decode(sale, r.PostForm)
		password := r.PostFormValue("password")
		if err = database.CompareAdminPassword(password); err != nil {
			utils.SetContext(r, "Success", nil)
			utils.SetContext(r, "Error", err)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")

			return
		}

		err = database.UpdateSaleById(id, sale)
		if err != nil {
			utils.SetContext(r, "Success", nil)
			utils.SetContext(r, "Error", err)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			return
		}

		utils.SetContext(r, "Success", "Se modificó el producto correctamente")
		utils.SetContext(r, "Error", nil)

	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func SalesQuery(w http.ResponseWriter, r *http.Request) {
	context := utils.GetFullContext(r)

	utils.RenderTemplate(w, "sales_query", context)
}
