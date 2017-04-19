package main

import (
	"log"
	"net/http"

	"github.com/mateuspontes/music-api/api"
)

func main() {
	addr := "127.0.0.1:8081"
	router := api.NewRouter()

	log.Printf("Hello music!")
	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
