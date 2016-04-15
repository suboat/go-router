package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const host_server string = "localhost:18080"

var upgrader = websocket.Upgrader{}

func logS(v ...interface{}) {
	log.Println("[S]:", fmt.Sprint(v...))
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logS("upgrade:", err)
		exit <- err
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			logS("read:", err)
			exit <- err
			break
		}
		logS(fmt.Sprintf("recv: %s", message))
		err = c.WriteMessage(mt, message)
		if err != nil {
			logS("write:", err)
			exit <- err
			break
		}
	}
}

func startServer() {
	http.HandleFunc("/echo", echo)
	logS("start...")
	exit <- http.ListenAndServe(host_server, nil)
}
