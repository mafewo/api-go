package main

import (
	"gomail/controllers/birthday"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/saracatunga", birthday.AllBirthdays).Methods("GET")
	myRouter.HandleFunc("/saracatunga", birthday.InsertBirthday).Methods("POST")
	return myRouter
}
