package main

import (
	"log"
	"net/http"
	"github.com/nihankhan/UserAuth-Go/handlers"
)

const address = ":8080"

func main() {
	// := db.Connect()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/login", handlers.LogIn)
	mux.HandleFunc("/q", handlers.Query)

	log.Println("Server listening on http://127.0.0.1" + address)
	log.Fatal(http.ListenAndServe(address, mux))	
}