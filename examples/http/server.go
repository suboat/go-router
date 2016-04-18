package main

import (
	"fmt"
	. "github.com/suboat/go-router/routers/http"
	. "github.com/suboat/go-router/routers/mux"
	"io"
	"log"
	"net/http"
)

const host_server string = "127.0.0.1:18081"

func logS(v ...interface{}) {
	log.Println("[S]:", fmt.Sprint(v...))
}

func startServer() {
	if true {
		r := NewHTTPRouter().PathPrefix("/v1")
		r.HandleFunc("/",
			func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, "hello 1")
			},
		).Methods("GET")
		r.HandleFunc("/id/*",
			func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, "hello 2")
			},
		).Methods("GET")
		logS("start...")
		exit <- r.Error()
		exit <- r.ListenAndServe(host_server)
	} else {
		r := NewMuxRouter().PathPrefix("/v1")
		r.HandleFunc("/",
			func(w http.ResponseWriter, r *http.Request) {
			},
		).Methods("GET")
		r.HandleFunc("/id/{id}",
			func(w http.ResponseWriter, r *http.Request) {
			},
		).Methods("GET")
		logS("start...")
		exit <- r.Error()
		exit <- r.ListenAndServe(host_server)
	}
}
