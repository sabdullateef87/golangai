package main

import (
	"fmt"
	"sync"
)

func main() {
	mut := sync.Mutex{}
	for i := 0; i < 100; i++ {
		fmt.Println("Total Items Packed:", PackItemsWithMutex(&mut, 0))
	}
}
func PackItems(totalItems int) int {
	const workers = 2
	const itemsPerWorker = 1000
	var wg sync.WaitGroup
	itemsPacked := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// Simulate the worker packing items into boxes.
			for j := 0; j < itemsPerWorker; j++ {
				itemsPacked = totalItems
				// Simulate packing an item.
				itemsPacked++
				// Update the total items packed without proper synchronization.
				totalItems = itemsPacked
			}
		}(i)
	}
	// Wait for all workers to finish.
	wg.Wait()
	return totalItems
}

func PackItemsWithMutex(m *sync.Mutex, totalItems int) int {
	const workers = 2
	const itemsPerWorker = 1000
	var wg sync.WaitGroup // this is used to wait for the go routines to finish before the main thread executes.
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done() // making this go routine register to the waitgroup. this process in this go-routine has to finish.
			for j := 0; j < itemsPerWorker; j++ {
				m.Lock() // this is to prevent data races, basically locking thr variable in memory 
				itemsPacked := totalItems
				itemsPacked++
				totalItems = itemsPacked
				m.Unlock() // this is use to unlock so that another process can access the variable.
			}
		}(i)
	}

	wg.Wait()
	return totalItems
}
