package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func client() {
	// connect to server
	c, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connection established")
	}

	// send message
	msg := "Hello World!"

	fmt.Println("Sending: ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Encoding ok")
	}

	// connection close
	err = c.Close()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connection closed")
	}
}

func main() {
	client()
}
