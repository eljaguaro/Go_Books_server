package main

import (
	"goserv/pkg/api"
	"goserv/pkg/repository"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	db, err := repository.New("SECRET")
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.Handle()
	// Нули для Докера
	log.Fatal(api.ListenAndServe("0.0.0.0:8090"))
}
