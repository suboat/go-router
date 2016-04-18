package main

import (
	"fmt"
	. "github.com/suboat/go-router/routers/gin"
	. "github.com/suboat/go-router/routers/mux"
	"log"
	"net/http"
)

const host_server string = "127.0.0.1:18081"

func logS(v ...interface{}) {
	log.Println("[S]:", fmt.Sprint(v...))
}

func startServer() {
	switch 1 {
	case 0:
		r := NewMuxRouter().PathPrefix("/v1")
		r.HandleFunc("/",
			func(w http.ResponseWriter, r *http.Request) {
				logS(fmt.Sprintf("%#v", r))
			},
		).Methods("GET")
		r.HandleFunc("/id/{id}",
			func(w http.ResponseWriter, r *http.Request) {
				logS(fmt.Sprintf("%#v", r))
			},
		).Methods("GET")
		logS("start...")
		exit <- r.Error()
		exit <- r.ListenAndServe(host_server)
	case 1:
		r := NewGinRouter().PathPrefix("/v1")
		r.HandleFunc("/",
			func(w http.ResponseWriter, r *http.Request) {
				logS(fmt.Sprintf("%#v", r))
			},
		).Methods("GET")
		r.HandleFunc("/id/:id",
			func(w http.ResponseWriter, r *http.Request) {
				logS(fmt.Sprintf("%#v", r))
			},
		).Methods("GET")
		logS("start...")
		exit <- r.Error()
		exit <- r.ListenAndServe(host_server)
	}
}
