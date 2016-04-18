package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	. "github.com/suboat/go-router/routers/mux"
	. "github.com/suboat/go-router/routers/websocket"
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

type MyUpgrader struct {
	Conn *websocket.Conn
}

func NewMyUpgrader() *MyUpgrader {
	return &MyUpgrader{}
}

func (s *MyUpgrader) Upgrade(w http.ResponseWriter, r *http.Request, h http.Header) (interface{}, error) {
	var err error
	upgrader = websocket.Upgrader{}
	s.Conn, err = upgrader.Upgrade(w, r, h)
	return s.Conn, err
}

func (s *MyUpgrader) ReadMessage() (int, []byte, error) {
	return s.Conn.ReadMessage()
}

func (s *MyUpgrader) WriteMessage(messageType int, data []byte) error {
	return s.Conn.WriteMessage(messageType, data)
}

func (s *MyUpgrader) Close() error {
	return s.Conn.Close()
}

func startServer() {
	switch 1 {
	case 0:
		exampleWebSocket()
	case 1:
		exampleWSRouter()
	}
}

func exampleWebSocket() {
	http.HandleFunc("/echo", echo)
	logS("start...0")
	exit <- http.ListenAndServe(host_server, nil)
}

func exampleWSRouter() {
	mux := NewMuxRouter().PathPrefix("/v1")
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			logS("Success!!!")
			logS(fmt.Sprintf("%#v", w))
			logS(fmt.Sprintf("%#v", r))
		},
	).Methods("GET")
	h, err := NewWSHandle(mux, NewMyUpgrader())
	if err != nil {
		exit <- err
	}
	r := NewWSRouter(h)
	r.Bind("/echo")
	logS("start...1")
	exit <- r.Error()
	exit <- r.ListenAndServe(host_server)
	exit <- r.Error()
}
