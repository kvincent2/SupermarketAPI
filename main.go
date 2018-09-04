package main

import (
	"log"
	"net/http"
	"github.com/kvincent2/SupermarketAPI/routes"
)

func main() {

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}


