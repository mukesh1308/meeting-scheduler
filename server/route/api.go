package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setApi(api *mux.Router){
	api.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("hello from api"))
	})
}

