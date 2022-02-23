package handlers

import (
	"fmt"
	"net/http"
	//"html/template"
)

// var (
// 	tmplt *template.Template
// )

// tmplt, err = template.ParseGlob("/home/nihan/Documents/UserAuth-Go/templates/*.html")

// if err != nil {
// 	panic(err)
// }


func Home(resp http.ResponseWriter, req *http.Request) {
	//tmplt.ExecuteTemplate(resp, "index.html", nil)

	fmt.Fprintf(resp, "Hello, Nihan Khan!!!\n")
}

func Test(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello, This is Test!!!\n")
}