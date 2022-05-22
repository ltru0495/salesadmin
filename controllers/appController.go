package controllers

import (
	"admin/models/database"
	"admin/utils"
	"log"
	"net/http"

	"admin/models"
	"time"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {

	user := utils.GetUser(r)
	if user.Role == "usr" {
		http.Redirect(w, r, "/venta/nueva", 301)
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

	if context == nil {
		http.Redirect(w, r, "/entrar", http.StatusSeeOther)
		return
	}

	if usr.Role == "admin" {
		sales, err := database.GetSalesByDate(t)

		sellers := models.GetSellers(sales)
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

		context["Title"] = "Ventas del día"
		context["TotalSalesC"] = totalSC
		context["TotalRefundsC"] = totalRC
		context["TotalEarningC"] = models.GetEarning(salesC)
		context["TotalSalesAndRefundsC"] = totalSC + totalRC

		context["TotalPEC"] = models.GetTotalSalesByPM(salesC, "electronico")
		context["TotalEC"] = models.GetTotalSalesByPM(salesC, "efectivo")
		context["UtilidadC"] = models.GetEarning(salesC)

		context["SalesC"] = salesC

		context["TotalSalesL"] = totalSL
		context["TotalRefundsL"] = totalRL
		context["TotalEarningL"] = models.GetEarning(salesL)
		context["TotalSalesAndRefundsL"] = totalSL + totalRL

		context["TotalPEL"] = models.GetTotalSalesByPM(salesL, "electronico")
		context["TotalEL"] = models.GetTotalSalesByPM(salesL, "efectivo")
		context["UtilidadL"] = models.GetEarning(salesL)

		context["SalesL"] = salesL

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
	}

	context["IsIndex"] = true
	utils.RenderTemplate(w, "index_admin", context)
	context["IsIndex"] = false

	utils.SetContext(r, "Error", nil)
	utils.SetContext(r, "Success", nil)
	utils.SetContext(r, "Sellers", nil)
	utils.SetContext(r, "Sales", nil)
}

/*
000002MA2302
000002MA2303
000003-BAN40CAJ01
000003MA2101
*/

func SalesList(w http.ResponseWriter, r *http.Request) {
	dates := database.GetAllDates()
	context := utils.GetFullContext(r)
	context["Dates"] = dates
	utils.RenderTemplate(w, "saleslist", context)
}

func BarcodesPrint(w http.ResponseWriter, r *http.Request) {
	context := utils.GetFullContext(r)
	utils.RenderTemplate(w, "product_barcodes_print", context)
}

func Charts(w http.ResponseWriter, r *http.Request) {
	context := utils.GetFullContext(r)
	utils.RenderTemplate(w, "charts", context)
}

func Users(w http.ResponseWriter, r *http.Request) {
	users := database.GetAllUsers()
	context := utils.GetFullContext(r)
	context["Users"] = users
	utils.RenderTemplate(w, "users", context)
	utils.SetContext(r, "Error", nil)
	utils.SetContext(r, "Success", nil)

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		utils.SetContext(r, "Error", err)
	} else {

		username := r.FormValue("username")
		oldPassword := r.FormValue("oldPassword")
		newPassword := r.FormValue("newPassword")

		err = database.UpdatePassword(username, oldPassword, newPassword)
		if err != nil {
			log.Println(err)
			utils.SetContext(r, "Error", err)
			return
		}

		utils.SetContext(r, "Success", "Se actualizó la contraseña")
	}

	http.Redirect(w, r, "/inventario", 301)
}

/*
Year 	06   2006
Month 	01   1   Jan   January
Day 	02   2   _2   (width two, right justified)
*/
