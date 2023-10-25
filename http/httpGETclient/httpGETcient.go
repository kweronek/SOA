package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Print("sent:\nHTTP GET:\n")
	res, err := http.Get("http://localhost:8080/"	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("received:")
	io.Copy(os.Stdout, res.Body)
}
