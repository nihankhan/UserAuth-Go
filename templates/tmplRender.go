package templates

import (
	"fmt"
	"log"
	"html/template"
)

func Render() (tmplt *template.Template, err error) {
	tmplt, err = template.ParseGlob("/home/nihan/Documents/UserAuth-Go/templates/*.html")

	fmt.Println("t : ", tmplt)

	if err != nil {
		log.Fatal("Unable to parse from template:- ", err)
	}

	return tmplt, err
}