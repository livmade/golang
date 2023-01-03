package main

import (
	"fmt"
)

func main() {
	var count = 10

	for count > 0 {
		fmt.Println("Launching in... ", count)
		count--
	}
	fmt.Println("Liftoff!")
}
