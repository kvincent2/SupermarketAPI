package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kvincent2/SupermarketAPI/produce"
	//"sync"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetProduce (w http.ResponseWriter, r *http.Request) {
	array := produce.Array
	if len(array) == 0 {
		log.Fatal("No produce available!")
	}
	arrayJSON, err := json.Marshal(array)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", arrayJSON)
	fmt.Fprintf(w, "%s", arrayJSON)
}

func GetProduceByID (w http.ResponseWriter, r *http.Request) {
	URLParams := r.URL.Query()
	produceCode := URLParams["ProduceCode"][0]
	array := produce.Array
	for _,v := range array {
		fmt.Println(v.ProduceCode, produceCode)
		if v.ProduceCode == produceCode {
			arrayJSON, err := json.Marshal(v)
			if err != nil {
				log.Fatal("Cannot encode to JSON ", err)
			}
			fmt.Fprintf(os.Stdout, "%s", arrayJSON)
			fmt.Fprintf(w, "%s", arrayJSON)
			return
		}
	}
	log.Fatal("Produce not found!")
}

func PostProduce (w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	var newItem produce.Produce
	//TODO Mutex Locking
	//c.mux.Lock()

	//TODO Figure out struct validation. https://github.com/asaskevich/govalidator; focus on 'validate struct' function.
	rules := govalidator.MapData{
		"Name": []string{"required","regex:^[0-9A-Za-z]$"},
		"ProduceCode":    []string{"required", "len:19", "regex:^([0-9A-Za-z]{4}-){3}[0-9A-Za-z]{4}$"},
		"UnitPrice":      []string{"required"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &newItem,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	fmt.Println(newItem) // your incoming JSON data in Go data struct

	validationErr := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "application/json")
	if validationErr != nil {
		//json.NewEncoder(w).Encode(validationErr)
		fmt.Println(validationErr)
	}
	if JSONErr := json.Unmarshal(rbody, &newItem); err != nil {
		http.Error(w, fmt.Sprintf("Can't read body. Error: %s",JSONErr), http.StatusBadRequest)
	}

	for _, item := range produce.Array {
		if strings.Contains(item.ProduceCode,newItem.ProduceCode) {
			http.Error(w, "Item already exists!", http.StatusBadRequest)
			return
		}

	}
	produce.Array = append(produce.Array,newItem)
	fmt.Print(produce.Array)
	fmt.Print("Success!")
}

func DeleteProduce (w http.ResponseWriter, r *http.Request) {
	deleteItem := r.URL.Query()
	produceCode := deleteItem["ProduceCode"][0]
	fmt.Println(produceCode)

	for k, v := range produce.Array {
		if v.ProduceCode == produceCode {
			produce.Array = append(produce.Array[:k], produce.Array[k+1:]...)
			fmt.Printf("Deleted ProduceCode : %+v", v)
			return
		}

	}
	log.Fatal("Produce not found!")
}


