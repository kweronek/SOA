package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"time"
)

func main() {

	var srv http.Server
	//Enable http2
	http2.ConfigureServer(&srv, nil)

	srv = http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", home)

	log.Fatal(srv.ListenAndServeTLS("certs/https-server.cert", "certs/https-server.key"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><center> Hello from Go using HTTP/2 ! </center></h1>"))
}
