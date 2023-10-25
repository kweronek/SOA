package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	const port = ":9999"
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server is running and listening on port", port, "\n")
	}

	// while true
	for {
		// accept a connection on listening port
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println("A client connection was accecpted")
		}

		// handle the server connection as go thread
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message and decode
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received:", msg)
	}

	// close connection
	err = c.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("A client connection was closed")
	}
}

func main() {
	go server()
	fmt.Println("Server started.\nPress enter to terminate.\n")
	var input string
	fmt.Scanln(&input)
}
