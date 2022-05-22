package controllers

import (
	"admin/models"
	"admin/models/database"
	"admin/utils"

	//"github.com/gorilla/mux"
	//"github.com/gorilla/schema"
	// "log"

	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		loginUser := models.User{
			Username: username,
			Password: password,
		}
		if user, err := database.Login(loginUser); err != nil {
			utils.SetContext(r, "Error", "Error de Autenticacion")
		} else {
			// log.Println(user)
			utils.SetSession(user, w)

			// dd-MM-yyyy

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	if utils.IsAuthenticated(r) {
		utils.SetContext(r, "Error", nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	context := utils.GetFullContext(r)
	utils.RenderTemplate(w, "login", context)
	utils.SetContext(r, "Error", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	utils.DeleteSession(w, r)
	http.Redirect(w, r, "/entrar", http.StatusSeeOther)
}
