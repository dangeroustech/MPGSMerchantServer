package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	port := os.Getenv("PORT")
	Logger("Starting MPGS Merchant Server on Port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
