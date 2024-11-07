package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipCode"`
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
	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
	}

}
