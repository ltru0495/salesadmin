package utils

import (
	"admin/config"
	"html/template"
	// "log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

var errorTemplate = template.Must(template.ParseFiles(config.DirTemplateError()))

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./templates/*.html",
	}
	rnd = renderer.New(opts)
}

func RenderErrorTemplate(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errorTemplate.Execute(w, nil)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	// log.Println(name)
	rnd.HTML(w, http.StatusOK, name, data)
}
