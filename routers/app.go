package routers

import (
	"admin/controllers"

	"github.com/gorilla/mux"
)

func SetAppRoutes(router *mux.Router) *mux.Router {
	appRouter := mux.NewRouter()

	appRouter.HandleFunc("/entrar", controllers.Login)
	appRouter.HandleFunc("/salir", controllers.Logout)

	appRouter.Handle("/", controllers.Authentication(controllers.IndexGET)).Methods("GET")
	appRouter.Handle("/graficos", controllers.Authentication(controllers.Charts)).Methods("GET")
	appRouter.Handle("/productos/imprimir", controllers.Authentication(controllers.BarcodesPrint)).Methods("GET")

	appRouter.Handle("/usuarios", controllers.HasPermission(controllers.Users)).Methods("GET")
	appRouter.Handle("/usuario/pass", controllers.HasPermission(controllers.UpdatePassword)).Methods("PUT")

	// GET REGISTER PRODUCT FORM AND POST PRODUCT
	appRouter.Handle("/productos", controllers.HasPermission(controllers.ManyProducts)).Methods("POST")
	appRouter.Handle("/producto/nuevo", controllers.HasPermission(controllers.NewProduct)).Methods("GET", "POST")
	appRouter.Handle("/producto/{code}", controllers.HasPermission(controllers.ViewProduct)).Methods("GET")
	appRouter.Handle("/producto/{code}", controllers.HasPermission(controllers.UpdateProduct)).Methods("PUT")
	appRouter.Handle("/producto/{code}", controllers.HasPermission(controllers.DeleteProduct)).Methods("DELETE")
	appRouter.Handle("/productos/reporte", controllers.HasPermission(controllers.InventoryReport)).Methods("GET")
	appRouter.Handle("/productos/reporte/validar", controllers.HasPermission(controllers.ValidateReport)).Methods("PUT")

	// INVENTORY TABLE GET
	appRouter.Handle("/inventario", controllers.Authentication(controllers.Inventory)).Methods("GET")

	// SALES
	appRouter.Handle("/venta/nueva", controllers.Authentication(controllers.NewSale)).Methods("GET", "POST")
	appRouter.Handle("/venta/{id}", controllers.HasPermission(controllers.ViewSale)).Methods("GET")
	appRouter.Handle("/venta/{id}", controllers.HasPermission(controllers.UpdateSale)).Methods("PUT")
	appRouter.Handle("/venta/{id}/eliminar", controllers.HasPermission(controllers.DeleteSale)).Methods("POST")
	appRouter.Handle("/venta/{id}/{action}", controllers.Authentication(controllers.RefundSale)).Methods("POST")

	appRouter.Handle("/ventas/devolver", controllers.Authentication(controllers.SalesRefund)).Methods("GET")
	appRouter.Handle("/ventas/buscar", controllers.Authentication(controllers.SalesQuery)).Methods("GET")
	appRouter.Handle("/ventas/{date}", controllers.Authentication(controllers.Sales)).Methods("GET")

	appRouter.Handle("/ventas/{date}/resumen", controllers.Authentication(controllers.GetSalesFile)).Methods("GET")
	appRouter.Handle("/ventas/{date}/resumen/pdf", controllers.Authentication(controllers.GetSalesPDFFile)).Methods("GET")
	appRouter.Handle("/ventasusuario/{date}/resumen", controllers.Authentication(controllers.GetSalesFileForUser)).Methods("GET")

	appRouter.Handle("/ventas/{place}/{date}", controllers.HasPermission(controllers.SalesByLocation)).Methods("GET")
	appRouter.Handle("/ventas/{place}/{date}/resumen", controllers.HasPermission(controllers.GetSalesFileByLocation)).Methods("GET")
	appRouter.Handle("/ventas/{place}/{date}/resumen/pdf", controllers.HasPermission(controllers.GetSalesPDFFileByLocation)).Methods("GET")

	appRouter.Handle("/ventas", controllers.Authentication(controllers.SalesList)).Methods("GET")

	// FILES
	appRouter.HandleFunc("/barcodes/pdf", controllers.GetBarcodesFile).Methods("GET")
	appRouter.Handle("/productos/ultimos", controllers.HasPermission(controllers.GetLastProductsFile)).Methods("GET")
	appRouter.Handle("/productos/faltan", controllers.HasPermission(controllers.GetUncheckedProductsFile)).Methods("GET")
	appRouter.HandleFunc("/productos/ultimos/barcodes/{type}", controllers.GetLastProductsBarcodesFile).Methods("GET")

	appRouter.Handle("/inventario/exportar", controllers.HasPermission(controllers.GetInventoryFile)).Methods("GET")

	router.PathPrefix("/").Handler(appRouter)

	return router

}
