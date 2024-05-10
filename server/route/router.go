package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.Handle("/", http.FileServer(http.Dir("./build/"))).Methods("GET")
	setApi(router.PathPrefix("/api").Subrouter())
	return router
}
