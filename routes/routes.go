package routes

import (
	// "net/http"

	"github.com/nihankhan/UserAuth-Go/mux"
	"github.com/nihankhan/UserAuth-Go/handlers"
)

func GetRoute() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home)

	return r
}

