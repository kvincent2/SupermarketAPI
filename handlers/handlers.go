package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/kvincent2/SupermarketAPI/produce"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
	arrayJSON, err := json.MarshalIndent(array,"	", "	")
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
		if v.ProduceCode == produceCode {
			arrayJSON, err := json.MarshalIndent(v, "		", "		")
			if err != nil {
				log.Fatal("Cannot encode to JSON ", err)
			}
			fmt.Fprintf(os.Stdout, "%s", arrayJSON)
			fmt.Fprintf(w, "%s", arrayJSON)
		} else {
			log.Fatal("Produce not found!")
		}

	}
}

func PostProduce (w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	var newItem produce.Produce
	//TODO Mutex Locking
	//c.mux.Lock()

	if err := json.Unmarshal(rbody, &newItem); err != nil {
		panic(err)
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
			produce.Array = append(produce.Array[:k],produce.Array[k+1:]...)
			fmt.Printf("Deleted ProduceCode : %s",v)
			break
		} else {
			log.Fatal("Produce not found!")
		}

	}
}


