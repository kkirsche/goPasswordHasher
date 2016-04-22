package main

import (
	"log"
	"net/http"
)

func main() {
	httpAddr := ":8080"
	httpsAddr := ":9090"

	http.HandleFunc("/public/", PublicAssetsHandler)
	http.HandleFunc("/hash", HashPasswordHandler)
	http.HandleFunc("/", CreateHashHandler)

	go func() {
		err := http.ListenAndServe(httpAddr, nil)
		if err != nil {
			log.Panicln(err)
		}
	}()

	err := http.ListenAndServeTLS(httpsAddr, "tls/cert.pem", "tls/key.pem", nil)
	if err != nil {
		log.Panicln(err)
	}
}
