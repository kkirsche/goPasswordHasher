package main

import (
	"log"
	"net/http"

	"github.com/kkirsche/ansiblePasswordGenerator/handlers"
)

func main() {
	http.HandleFunc("/public/", hasherHandlers.PublicAssets)
	http.HandleFunc("/hash", hasherHandlers.HashPasswordHandler)
	http.HandleFunc("/", hasherHandlers.CreateHashHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
