package main

import (
	"net/http"
)

//Route -
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes -
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Auth(Index),
	},
	Route{
		"StartPayment",
		"POST",
		"/startpayment",
		Auth(StartPayment),
	},
	Route{
		"FinishPayment",
		"PUT",
		"/finishpayment",
		Auth(FinishPayment),
	},
}
