package main

import (
	"log"
	"os"
	"time"
)

var exit = make(chan error)

func main() {
	go startServer()
	<-time.After(time.Second)
	go startClient()
	select {
	case <-time.After(time.Minute):
		os.Exit(0)
	case err := <-exit:
		log.Fatal(err)
	}
}
