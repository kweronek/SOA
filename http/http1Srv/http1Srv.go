package main

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

	http.HandleFunc("/country", country)

	srv.ListenAndServe()
}

func country(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://country.io")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("received:", time.Now())
	io.Copy(w, res.Body)
}
