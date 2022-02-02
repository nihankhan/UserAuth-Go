package main

import (
	"log"
	"net/http"
	"github.com/nihankhan/UserAuth-Go/routes"
	"github.com/nihankhan/UserAuth-Go/handlers"
)

const address = ":8080"

func main() {
	route := routes.GetRoutes()

	log.Println("Server listening on http://127.0.0.1" + address)
	log.Fatal(http.ListenAndServe(address, mux))	
}