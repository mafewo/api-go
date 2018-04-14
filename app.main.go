package main

import (
	"log"
	"net/http"
)

func main() {
	router := getRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
