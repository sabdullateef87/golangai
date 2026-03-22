package main

import (
	"fmt"
	"os"
	"time"
)

func main () {
	fmt.Println ("Understanding select in golang")

	ch := make(chan string)

	// inline g0-routine (sends data into the channel)
	go func () {
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("Hello")
	} ()

	//Another inline g0-routine (sends data into the channel)
	go func () {
		time.Sleep(5 * time.Second)
		ch <- fmt.Sprintf("world")
	} ()


	for {

		// Listens to whoever is ready from any number of channels and handle then different based on the business needs.
		select {
			case v := <- ch:
				fmt.Printf("%s\n", v)
			case <- time.After(3 * time.Second):
				fmt.Println("I have waited for 3 seconds and no one was ready!.")
				fmt.Println("Quitting this process")
				os.Exit(0)
			}
	}
}