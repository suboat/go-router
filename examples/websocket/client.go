package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	. "github.com/suboat/go-router/routers/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func logC(v ...interface{}) {
	log.Println("[C]:", fmt.Sprint(v...))
}

func startClient() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: host_server, Path: "/echo"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		exit <- err
		return
	}
	defer c.Close()
	done := make(chan struct{})
	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logC("read:", err)
				exit <- err
				return
			}
			logC(fmt.Sprintf("recv: %s", message))
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			if b, err := makeWSRequest(t).Bytes(); err != nil {
				exit <- err
			} else if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
				logC("write:", err)
				exit <- err
			}
		case <-interrupt:
			logC("interrupt")
			if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				logC("write close:", err)
				exit <- err
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}

func makeWSRequest(data interface{}) *WSRequest {
	return &WSRequest{
		Tag: "Tag",
	}
}
