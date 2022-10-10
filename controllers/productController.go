package controllers

import (
	"admin/models"
	"admin/models/database"
	"admin/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func ManyProducts(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	var prods []models.Product
	err = json.Unmarshal([]byte(r.PostFormValue("data")), &prods)
	if err != nil {
		log.Println(err)
	}
	code := ""
	quantity := 1
	for _, product := range prods {
		code = product.Code
		quantity = product.Quantity
		for k := 1; k <= quantity; k++ {
			product.Code = code + fmt.Sprintf("%03d", k)
			product.Quantity = 1
			log.Println(product)
			err = database.InsertProduct(&product)
			if err != nil {
				break
			}

		}
		if err != nil {
			log.Println(err)
			utils.SetContext(r, "Error", "Ocurrio un error")
			return

		} else {
			utils.SetContext(r, "Success", "Se guardó el producto correctamente")

			brand := &models.Brand{Name: product.Brand}
			model := &models.Model{Name: product.Model}

			err = database.InsertBrand(brand)
			if err != nil {
				log.Println(err)
			}

			err = database.InsertModel(model)
			if err != nil {
				log.Println(err)
			}
		}
	}

	err = database.UpdateCounter()
	if err != nil {
		log.Println(err)
	}

}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			// utils.SetContext(r, "Error", "Ocurrio un error")
		} else {
			product := new(models.Product)
			// GET from form

			decoder := schema.NewDecoder()

			err = decoder.Decode(product, r.PostForm)

			if err != nil {
				log.Println(err)
				utils.SetContext(r, "Error", "Ocurrio un error")
			} else {
				//Save

				err = database.InsertProduct(product)
				if err != nil {
					log.Println(err)
					utils.SetContext(r, "Error", "Ocurrio un error")
					return

				} else {
					utils.SetContext(r, "Success", "Se guardó el producto correctamente")

					brand := &models.Brand{Name: product.Brand}
					model := &models.Model{Name: product.Model}

					err = database.InsertBrand(brand)
					if err != nil {
						log.Println(err)
					}

					err = database.InsertModel(model)
					if err != nil {
						log.Println(err)
					}
				}
			}
			http.Redirect(w, r, "/producto/nuevo", 301)

		}
	}
	context := utils.GetFullContext(r)
	counter, err := database.GetCounter()
	if err != nil {
		log.Println(err)
	}
	context["Counter"] = fmt.Sprintf("%06d", (counter.Counter + 1))

	utils.RenderTemplate(w, "product_new", context)
	utils.SetContext(r, "Error", nil)
	utils.SetContext(r, "Success", nil)

}

func ViewProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	product, err := database.GetProductByCode(code)
	if err != nil {
		log.Println(err)
		return
	}
	context := utils.GetFullContext(r)
	context["Product"] = product

	utils.RenderTemplate(w, "product_view", context)
	utils.SetContext(r, "Success", nil)
	utils.SetContext(r, "Error", nil)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		utils.SetContext(r, "Error", err)
	} else {
		product := new(models.Product)
		decoder := schema.NewDecoder()
		err = decoder.Decode(product, r.PostForm)
		// password := r.PostFormValue("password")
		// if err = database.CompareAdminPassword(password); err != nil {
		// 	utils.SetContext(r, "Success", nil)
		// 	utils.SetContext(r, "Error", err)
		// 	w.WriteHeader(http.StatusForbidden)
		// 	w.Header().Set("Content-Type", "application/json")

		// 	return
		// }

		err = database.UpdateProductByCode(code, product)
		if err != nil {
			utils.SetContext(r, "Success", nil)
			utils.SetContext(r, "Error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			return
		}

		utils.SetContext(r, "Success", "Se modificó el producto correctamente")
		utils.SetContext(r, "Error", nil)

	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]

	err := database.DeleteProductByCode(code)
	if err != nil {
		// utils.SetContext(r, "Error", err)
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	res := models.CreateDefaultResponse(w)

	res.Send()
}

func Inventory(w http.ResponseWriter, r *http.Request) {
	context := utils.GetFullContext(r)
	user := utils.GetUser(r)
	if user.Role == "admin" {
		utils.RenderTemplate(w, "inventory", context)
	} else {
		utils.RenderTemplate(w, "inventory_user", context)
	}
	utils.SetContext(r, "Error", nil)
	utils.SetContext(r, "Success", nil)
	utils.SetContext(r, "Products", nil)
}

func InventoryReport(w http.ResponseWriter, r *http.Request) {
	context := utils.GetFullContext(r)
	products := database.GetAllProducts()
	context["Products"] = products

	utils.RenderTemplate(w, "report", context)

	utils.SetContext(r, "Error", nil)
	utils.SetContext(r, "Success", nil)
	utils.SetContext(r, "Products", nil)
}

func ValidateReport(w http.ResponseWriter, r *http.Request) {
	if err := database.SetAllProductsUnchecked(); err != nil {
		utils.SetContext(r, "Success", nil)
		utils.SetContext(r, "Error", err)
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	utils.SetContext(r, "Success", "Se validaron los productos correctamente")
	utils.SetContext(r, "Error", nil)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func LastProductsBarcodes(w http.ResponseWriter, r *http.Request) {
	context := utils.GetFullContext(r)
	utils.RenderTemplate(w, "lastbarcodes", context)

}

/*
Year 	06   2006
Month 	01   1   Jan   January
Day 	02   2   _2   (width two, right justified)
*/
