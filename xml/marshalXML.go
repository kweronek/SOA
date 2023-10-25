package main

import (
	"encoding/xml"
	"fmt"
)

type Message struct {
	Name string		`xml:"Name"`
	Body string		`xml:"Body"`
	Time int64		`xml:"Time"`
}

func Encode() {
	m := Message{"Alice",
		"Hello",
		1294706395881547000}
	fmt.Println("\nEncode:\n", m)

	j, err := xml.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded XML-Object:\n%s%s", j, "\n\n")
	fmt.Println("Alternative:")
	pretty, err := xml.MarshalIndent(m, " ", "    ")
	fmt.Println(string(pretty))
}

func Decode() {
	x := "<Message><Name>Alice</Name><Body>Hello</Body><Time>1294706395881547000</Time></Message>"
	fmt.Println("\nDecode:\n", x)

	var m Message
	err := xml.Unmarshal([]byte(x), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded string:\n%+v%s", m, "\n\n")
}

func PartialDecode() {
	b := []byte(`<Message><Name>Bob</Name><Food>Pickle</Food></Message>`)

	var m Message
	err := xml.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	expected := Message{
		Name: "Bob",
	}
	fmt.Printf("m:        %+v%s", m, "\n")
	fmt.Printf("expected: %+v%s", expected, "\n")
}

func main() {
	Encode()
	Decode()
	PartialDecode()
}