
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	n, _ := strconv.ParseInt(os.Args[1], 10, 0)
	for i := -n+1; i < n; i++ {
		for j := -2*n+2; j <= 2*n-1; j++ {
			switch {
				case j == 2*n-1:
					fmt.Println()
				case Abs(j/2) + Abs(i) < n && j%2 == 0:
					fmt.Printf("%c", Abs(j/2) + Abs(i) + 97)
				default :
					fmt.Printf("%c",45)
			}
		}
	}
}




func Abs(i int64) int64 {
	if i < 0{
		return -i}
		return i
}
