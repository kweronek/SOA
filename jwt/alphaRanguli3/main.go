package main

import (
	"fmt"
	"os"
	"strconv"
)

const 	Alphabet string = "f-e-d-c-b-a-b-cdefghijklmnopqrstuvwxyz"


func main(){
	n, _ := strconv.ParseInt(os.Args[1], 10, 0)

	for i := -n+1; i < n; i++ {
		for j := -2*n+2; j < 2*n-1; j++ {
			if Abs(j/2) + Abs(i) >= n || j%2 != 0 {
				fmt.Printf("%s","-")
			} else {
				fmt.Printf("%c", Abs(j/2) + Abs(i) + 97)
			}
		}
		fmt.Println()
	}
}

func Abs(i int64) int64 {
	if i < 0{
		return -i}
	return i
}
