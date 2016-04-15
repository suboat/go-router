package main

import (
	"fmt"
	"github.com/suboat/go-router/routers/mux"
	"log"
	"net/http"
)

const host_server string = "127.0.0.1:18081"

func logS(v ...interface{}) {
	log.Println("[S]:", fmt.Sprint(v...))
}

func startServer() {
	r := mux.NewMuxRouter().PathPrefix("/v1")
	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
		},
	).Methods("GET")
	r.HandleFunc("/id/{id}",
		func(w http.ResponseWriter, r *http.Request) {
		},
	).Methods("GET")
	logS("start...")
	exit <- r.ListenAndServe(host_server)
}
