package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		log.Printf("reqBody=%s err=%v", reqBody, err)
	})
	http.ListenAndServe(":8080",nil)
}
