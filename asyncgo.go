package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Printf("Understanding asynchronous Go routines\n")
	c := make(chan int)
	go PrintNumbers(c)
	fmt.Println("......................")

	// Reading the results of the execution of the go routine.
	for value := range c {
		switch value {
			case 0:
				fmt.Println("Received O")
			
			case 1:
				fmt.Println("Received 1")
			
			default:
				fmt.Printf("Received other values %d\n", value)
			}
	}
}

func PrintNumbers(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i // feeding data into a channels
		time.Sleep(100 * time.Millisecond)
	}
	close(c) // closing tht channel when the operation is done.
}


// Basically channels acts more like sequential in webflux
// gets all the result of the go routine and consolidate it in the channel, then
// whoever needs it can get the result as an array item.
