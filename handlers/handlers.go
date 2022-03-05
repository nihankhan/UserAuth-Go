package handlers

import (
	"fmt"
	"net/http"
	// "html/template"

	"github.com/nihankhan/UserAuth-Go/db"
	"github.com/nihankhan/UserAuth-Go/templates"
)

type nihan struct {
	ID       int
	fullName string
	username string
	password string
}

func Index(resp http.ResponseWriter, req *http.Request) {
	tmpl, _ := templates.Render()

	//fmt.Println(tmpl)

	tmpl.ExecuteTemplate(resp, "index.html", nil)
}

func LogIn(resp http.ResponseWriter, req *http.Request) {
	//tmpl, _ := template.ParseGlob("/home/nihan/Documents/UserAuth-Go/templates/*.html")

	tmpl, _ := templates.Render()

	tmpl.ExecuteTemplate(resp, "login.html", nil)
}

func Query(resp http.ResponseWriter, req *http.Request) {

	// Connect database
	db := db.Connect()

	qry, err := db.Query("SELECT * FROM Nihan.users")

	if err != nil {
		panic(err)
	}

	defer qry.Close()

	for qry.Next() {
		var ebu nihan

		err := qry.Scan(&ebu.ID, &ebu.fullName, &ebu.username, &ebu.password)

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(resp, "%v\n", ebu)
	}
}