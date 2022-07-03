package routers

import (
	"admin/controllers/api"

	"github.com/gorilla/mux"
)

func SetApiRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/inventory/{page}", api.Inventory).Methods("GET")

	router.HandleFunc("/api/producto/reporte/{code}", api.ProductChecked).Methods("POST")

	router.HandleFunc("/api/venta/{code}", api.GetSale).Methods("GET")

	router.HandleFunc("/api/ventas/{location}/{start}/{end}", api.QuerySales).Methods("GET")
	router.HandleFunc("/api/ventas/exportar/{location}/{start}/{end}", api.GetSalesFile).Methods("GET")

	router.HandleFunc("/api/producto/{code}", api.GetProduct).Methods("GET")
	router.HandleFunc("/api/producto/{code}", api.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/marcas", api.GetBrands).Methods("GET")
	router.HandleFunc("/api/modelos", api.GetModels).Methods("GET")
	router.HandleFunc("/api/barcodes", api.GetLastProducts).Methods("GET")

	router.HandleFunc("/api/graficos/marcas", api.GetBrandsForChart).Methods("GET")
	router.HandleFunc("/api/graficos/series", api.GetSeriesForChart).Methods("GET")

	router.HandleFunc("/api/productos/barcodes", api.ProductsBarcodeFile).Methods("POST")
	router.HandleFunc("/api/productos/barcodes2", api.ProductsBarcodeFile2).Methods("POST")

	return router
}
