package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int `json:"ref"`
	private string // An unexported field is not encoded.
	Number  int
	Created time.Time
}

func main() {

	//  Part 1 ********************************************************************
	fmt.Println("\nPart 1 ****************************************\n")
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
		Number:  55,
		Created: time.Now(),
	}

	var jsonData []byte

	// Encode (plain) / Marshal
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	fmt.Println()

	// Encode (pretty print) / Marshal
	jsonData, err = json.MarshalIndent(basket, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	fmt.Println()

	// Decode / Unmarshal
	err = json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id, basket.Number)
	fmt.Println(basket.Created)
	fmt.Println()

	//  Part 2 ********************************************************************

	fmt.Println("Part 2 ****************************************\n")
	jsonData2 := []byte(`{"Name":"Eve","Age":6 ,"Parents":["Alice","Bob"], "ok":true, "Score":1.41}`)

	var v interface{}
	json.Unmarshal(jsonData2, &v)
	data := v.(map[string]interface{})

	for k, v := range data {  // v entspricht data[k]

//		fmt.Println(k, v , "type: ",reflect.TypeOf(data[k]))
		switch v := v.(type) {
		case int:
			fmt.Println(k, ":", v, "-> type: (int64)")
		case string:
			fmt.Println(k, ":", v, "-> type: (string)")
		case []interface{}:
			fmt.Println(k, ":", "-> type: (array):")
			for i, n := range v {
				fmt.Println(k, "    ", i, ":", n)
			}
		case bool:
			fmt.Println(k, ":", v, "-> type: (bool)")
		case float64:
			{
				val1 := v
				val2 := int(val1)
				if val1 == float64(val2) {
					fmt.Println(k, v, "-> type: (", reflect.TypeOf(val2),")") // 10, int
				} else {
				fmt.Println(k, ":", v, "-> type: (float64)")
				}
			}
		default:
			fmt.Println(k, ":", v, "-> type: (unknown)")
		}
	}

	//  Part 3 ********************************************************************

	fmt.Println("\nPart 3 ****************************************\n")

}
