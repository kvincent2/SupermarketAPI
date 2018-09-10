package main

import (
	"fmt"
	"github.com/kvincent2/SupermarketAPI/routes"
	"net/http"
	"os"
)

func main() {

	router := routes.NewRouter()
	fmt.Println("Serving on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error: %s", err)
		fmt.Fprintf(os.Stdout, "Error: %s", err)
	}
}


