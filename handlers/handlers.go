package handlers

import (
	"net/http"
	"html/template"
)

// var (
// 	Tmplt *template.Template
// )

Tmplt, err = template.ParseGlob("/home/nihan/Documents/UserAuth-Go/templates/*.html")

if err != nil {
	panic(err)
}


func Home(resp http.ResponseWriter, req *http.Request) {
	Tmplt.ExecuteTemplate(resp, "index.html", nil)
}