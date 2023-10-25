package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	srv := http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", home)

	log.Fatal(srv.ListenAndServeTLS("certs/https-server.cert", "certs/https-server.key"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><center> Hello from Go using HTTP1.1 with TLS ! </center></h1>"))
}
