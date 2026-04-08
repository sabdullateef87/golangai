package main

import (
	"fmt"
	"sync"
	"time"
)

// Channels are used to let goroutines communicate safely by sending values between them and coordinating concurrent work.
func main() {
	fmt.Println("Looking at buffered channels")

	clownChannel := make(chan int, 3)
	clowns := 5

	var wg sync.WaitGroup

	// Here is the sender logic
	for clown := 1; clown <= clowns; clown++ {
		wg.Add(1) // Addding the below inline go function into the waitgroup
		go func(clownID int) {
			defer wg.Done() // calling defer early in order for the waigroup and also marking this go routine as a participant in the wait group.

			balloon := fmt.Sprintf("Balloon %d", clownID)
			fmt.Printf("Clown %d: Hopped into the car with %s\n", clownID, balloon)
			select {
			case clownChannel <- clownID: // sending into the channel (clownchannel)
				fmt.Printf("Clown %d: Finished with %s\n", clownID, balloon)
			default:
				fmt.Printf("Clown %d: Oops, the car is full, can't fit %s!\n", clownID, balloon)
			}
		}(clown)
	}

	go func() {
		defer close(clownChannel)
		for clownID := range clownChannel {
			balloon := fmt.Sprintf("Balloon %d", clownID)
			fmt.Printf("Driver: Drove the car with %s inside\n", balloon)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Driver: Clown finished with %s, the car is ready for more!\n", balloon)
		}
	}()

	wg.Wait() // this is used to wait for all the go routines in the waitgroup to complete execution before the main method terminates.
}
