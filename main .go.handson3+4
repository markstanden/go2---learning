package main

import (
	"fmt"
	"sync"
)

var glCounter int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	for i := 0; i < 100; i++ {

		count(&glCounter)

	}

	wg.Wait()
	fmt.Println("Counter :", glCounter)
}

func count(c *int) {
	wg.Add(1)
	mu.Lock()

	glCounter = (*c + 1)

	mu.Unlock()
	wg.Done()
}
