package handlers

import (
	"net/http"
	"html/template"
)

var (
	tmplt *template.Template
)

tmplt, _ = template.ParseGlob("/home/nihan/Documents/UserAuth-Go/templates/*.html")


func Home(resp http.ResponseWriter, req *http.Request) {
	tmplt.ExecuteTemplate(resp, "index.html", nil)
}