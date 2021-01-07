package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	for {
		request()
		time.Sleep(time.Millisecond * 10)
	}
}

func request() {
	client := http.Client{}
	body := strings.NewReader("132")
	req, err := http.NewRequestWithContext(context.Background(), "POST", "http://127.0.0.1:8080/", body)
	if err != nil {
		log.Printf("NewRequestWithContext err=%v", err)
		return
	}
	rsp, err := client.Do(req)
	if err != nil {
		log.Printf("Do err=%v", err)
		return
	}
	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("ReadAll err=%v", err)
		return
	}
	log.Printf("%s", data)
}
