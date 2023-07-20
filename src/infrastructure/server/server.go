package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"url-shortener/src/links"
)

func RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/short", links.GetShortened).Methods("POST")
	router.HandleFunc("/origin", links.GetOrigin).Methods("POST")
	//router.HandleFunc("/{shortURL}", links.UrlRedirect)

	err := http.ListenAndServe(":8090", router) // todo: use config
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}

}
