package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	http.ListenAndServe("localhost:8000", nil)
}

func greet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")
}

func getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"Ali", "mashhad", "1234"},
		{"Amin", "mashhad", "12345"},
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(customers)
}
