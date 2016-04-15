package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func logC(v ...interface{}) {
	log.Println("[C]:", fmt.Sprint(v...))
}

func getUrl(uri string) string {
	return "http://" + host_server + uri
}

func startClient() {
	doRequest(http.NewRequest("GET", getUrl("/v1/"), nil))
	doRequest(http.NewRequest("GET", getUrl("/v1/id/ABCabc123"), nil))

	os.Exit(0)
}

func doRequest(req *http.Request, err error) {
	if err != nil {
		exit <- err
	} else if resp, err := http.DefaultClient.Do(req); err != nil {
		exit <- err
	} else {
		logC(fmt.Sprintf("%#v", resp))
	}
}
