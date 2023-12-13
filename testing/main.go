package main

import (
    "fmt"
)

func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main () {
    fmt.Println(IntMin(3,6))
    fmt.Println(IntMin(-3,-6))
}
