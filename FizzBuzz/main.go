package main

// FizzBuzz //
/*
Print the numbers between 1 to 20, BUT!
- If the number is divisble by 3, print "Fizz"
- If the number is divisble by 5. print "Buzz"
*/

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {

	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			//If 'i' is equal to something divisible by 5
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			//If 'i' is equal to something divisible by 5
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			//If 'i' is equal to something divisible by 3
			fmt.Println("Fizz")
		} else {
			fmt.Printf("%d \n", i)
		}

	}

}
