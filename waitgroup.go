package main

import (
	"fmt"
	"sync"
)

const (
	workers = 2
	itemsPerWorker = 1000
)
func main () {
	fmt.Print("Total Items Packed : ", PackItems(0))
}

func PackItems (totalItems int) int {
	
	var wg sync.WaitGroup
	wg.Add(workers)
	itemsPacked := 0
	for i := 0; i < workers; i++ {
		// wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				itemsPacked ++
				totalItems = itemsPacked
			}
		}(i)
	}
	wg.Wait()
	return itemsPacked
}