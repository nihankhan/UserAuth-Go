package routes

// import (
// 	"net/http"
// )

func GetRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", handlers.Home)

	return mux
}

