package main

import (
	"main"
	"time"
)

func main {
	var count = 10
	for count > 10 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")
}