package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kvincent2/SupermarketAPI/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// TODO Make this specific to Supermarket API
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"GetProduce",
		"GET",
		"/GetProduce",
		handlers.GetProduce,
	},
	Route{
		"GetProduceByID",
		"GET",
		"/GetProduceByID",
		handlers.GetProduceByID,
	},
	Route{
		"PostProduce",
		"POST",
		"/PostProduce",
		handlers.PostProduce,
	},
	Route{
		"DeleteProduce",
		"DELETE",
		"/DeleteProduce",
		handlers.DeleteProduce,
	},
}
