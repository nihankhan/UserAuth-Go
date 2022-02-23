package main

import (
	"log"
	"net/http"
	"github.com/nihankhan/UserAuth-Go/db"
	"github.com/nihankhan/UserAuth-Go/handlers"
)

const address = ":8082"

func main() {
	db.Connect()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/test", handlers.Test)

	log.Println("Server listening on http://127.0.0.1" + address)
	log.Fatal(http.ListenAndServe(address, mux))	
}