package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	type Address struct {
		Street string
		City   string
		State  string
		Zip    string
	}
	type Person struct {
		Name    string
		Address Address
	}

	data := `
    name: John Doe
    address: 
        street: 123 E 3rd St
        city: Denver
        state: CO
        zip: 81526
    `

	var person Person
	err := yaml.Unmarshal([]byte(data), &person)
	if err != nil {
		fmt.Printf("Failed to unmarshall: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Marshalled person=%v", person)
}
