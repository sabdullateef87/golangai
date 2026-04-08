package main

import (
	"fmt"
	"sync"
)

func main() {
	balls := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go throwBalls("red", balls, &wg)
	go throwBalls("green", balls, &wg)
	go func() {
		wg.Wait()
		close(balls)
	}()
	
	for val := range balls {
		fmt.Println(val)
	}

	wg.Wait()
}
func throwBalls(color string, balls chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("throwing the %s ball\n", color)
	for i := range 5 {
		balls <- fmt.Sprintf("Processing ball number :  %d with color %s", i, color)
	}
}
