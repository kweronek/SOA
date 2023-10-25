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

	//	expected := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	//	if !reflect.DeepEqual(b, expected) {
	//		log.Panicf("Error marshaling %q, expected %q, got %q.", m, expected, b)
	//	}
	fmt.Printf("Encoded YAML-Object:\n%s\n\n", j)
}

func Decode() {

	y := `
    name: Alice
    body: Hello
    time: 1294706395881547000`

	fmt.Println("Decode:", string(y))

	var m Message
	err := yaml.Unmarshal([]byte(y), &m)
	if err != nil {
		panic(err)
	}

	/*
		expected := Message{
			Name: "Alice",
			Body: "Hello",
			Time: 1294706395881547000,
		}

		if !reflect.DeepEqual(m, expected) {
			log.Panicf("Error unmarshaling %q, expected %q, got %q.", b, expected, m)
		}
	*/
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

	//	if !reflect.DeepEqual(expected, m) {
	//		log.Printf("Error unmarshaling:\n%q\nexpected: %q\ngot:      %q", b, expected, m)
	//	}
}

func main() {
	Encode()
	Decode()
	PartialDecode()
}
