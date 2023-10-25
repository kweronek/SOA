package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id      string
	PIN		uint64
}

func main() {
	myStruct := User{Id: "SOA-Class", PIN: 4711}
	req := new(bytes.Buffer)
	json.NewEncoder(req).Encode(myStruct)

	// Alternative:
	//	e := json.NewEncoder(req)
	//	e.Encode(myStruct)

	fmt.Print("sent:\nPOST: ")
	io.Copy(os.Stdout, req)
	res, err := http.Post("http://localhost:8080/resources", "application/json; charset=utf-8", req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("received:")
	io.Copy(os.Stdout, res.Body)
}
