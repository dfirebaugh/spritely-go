package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir(".dist"))
	http.Handle("/", fs)

	log.Println("Listening on :3002...")
	err := http.ListenAndServe(":3002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
