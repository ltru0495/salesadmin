package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	router = SetApiRoutes(router)

	router = SetAppRoutes(router)

	return router
}
