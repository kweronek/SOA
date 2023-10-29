package main

import ("net/http")

func main() {http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)}

func home(w http.ResponseWriter, r *http.Request) {w.Write([]byte("<h1><center> Hello from Go using HTTP/1.1 ! </center></h1>"))}
