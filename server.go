package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(response http.ResponseWriter, request *http.Request) {
	time.Sleep(5 * time.Second)
	response.WriteHeader(200)
	response.Write([]byte("OK"))
}
