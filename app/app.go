package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	//router := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
