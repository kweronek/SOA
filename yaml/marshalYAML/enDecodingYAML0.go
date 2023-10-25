package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Message struct {
	Name string `yaml:"name,omitempty"`
	Body string `yaml:"body,omitempty"`
	Time string `yaml:"time,omitempty"`
}

func Encode() {
	m := Message{"Alice",
		"Hello",
		"1294706395881547000"}
	fmt.Println("\nEncode:\n", m)

	j, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded YAML-Object:\n%s\n\n", j)
}

func Decode() {

	y := `
    name: Alice
    body: Hello
    time: 1294706395881547000
    `

	fmt.Println("Decode:", string(y))

	var m Message
	err := yaml.Unmarshal([]byte(y), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nDecoded string:\n%s\n\n", m)
}

func PartialDecode() {
	y := `
    name: Bob
    food: Pickle
    `

	var m Message
	err := yaml.Unmarshal([]byte(y), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println("To decode:", y)

	expected := Message{
		Name: "Bob",
	}
	fmt.Println("Partially decoded:")
	fmt.Printf("m:        %v\n", m)
	fmt.Printf("expected: %v\n\n", expected)
}

func main() {
	Encode()
	Decode()
	PartialDecode()
}
