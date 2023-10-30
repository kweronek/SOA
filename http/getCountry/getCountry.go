package main
// einfacher http-Server, konfiguriert als Umleitung auf die Seite http://country.io

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/capital", capital)
	http.HandleFunc("/", country)

	srv.ListenAndServe()
}

func capital(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://country.io/capital.json"	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("received:")
	io.Copy(w, res.Body)
}

func country(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://country.io"	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("received:")
	io.Copy(w, res.Body)
}
