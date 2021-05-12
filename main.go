package main

import (
	"log"
	"net/http"

	"github.com/lonysutrisno/tokenomy/pkg"
	"github.com/lonysutrisno/tokenomy/service/checkout/delivery"
)

func main() {
	router := delivery.NewRouter()
	log.Println("ListenAndServe tokenomy-svc:8082")
	err := http.ListenAndServe(":8082", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
