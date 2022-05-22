package controllers

import (
	"admin/utils"
	// "fmt"
	"net/http"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

func Authentication(function customeHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context := utils.GetFullContext(r)

		if !utils.IsAuthenticated(r) || context == nil {
			utils.DeleteSession(w, r)
			http.Redirect(w, r, "/entrar", http.StatusSeeOther)
			return
		}

		function(w, r)
	})
}

func HasPermission(function http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context := utils.GetFullContext(r)
		if !utils.IsAuthenticated(r) || (utils.IsAuthenticated(r) && context == nil) {
			http.Redirect(w, r, "/entrar", http.StatusSeeOther)
			return
		}
		if !utils.IsAdmin(r) {
			http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
			return
		}
		function(w, r)
	})
}
