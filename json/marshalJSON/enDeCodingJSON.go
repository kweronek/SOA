package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string		`json:"name"`
	Body string     `json:"body"`
	Time int64      `json:"time"`
}

func Encode() {
	m := Message{"Alice",
		"Hello",
		1294706395881547000}
	fmt.Println("\nEncode:\n", m)

	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded JSON-Object:\n%s\n\n", j)
	fmt.Println("Alternative:")
	pretty, err := json.MarshalIndent(m,"","   ")
	fmt.Println(string(pretty))
}

func Decode() {
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	fmt.Println("\nDecode:\n", string(b))

	var m Message
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded string:\n%+v\n\n", m)
}

func PartialDecode() {
	b := []byte(`{"name":"Bob","food":"Pickle"}`)

	var m Message
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	expected := Message{
		Name: "Bob",
	}
	fmt.Printf("Partial Decode:\n %s\n", string(b))
	fmt.Printf("expected: %+v\n", expected)
	fmt.Printf("m:        %+v\n", m)
}

func main() {
	Encode()
	Decode()
	PartialDecode()
}