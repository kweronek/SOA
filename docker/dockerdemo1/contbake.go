package main

import (
	"fmt"
	"os"
	"time"
)

// Main Program
func main() {
	// init
	max := 10
	tStart := time.Now().Format(time.StampMilli)
	fmt.Println("Container-Bake gestartet!")

	// loop
	for i := 1; i <= max; i++ {

		fmt.Printf(
			"Tag: %3d /%3d | PID: %d | stamp: %v | start: %v\n",
			i,
			max,
			os.Getpid(),
			time.Now().Format(time.StampMilli),
			tStart)

		time.Sleep(10 * time.Second)
	}
	// end
	fmt.Println("Container-Bake beendet!")
}
