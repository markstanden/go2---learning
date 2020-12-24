package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		
		go count(&counter)
		
		wg.Done()
	}

	wg.Wait()
	fmt.Println("Counter :", counter)
}

func count(cPoint *int) {

	mu.Lock()
	c := *cPoint
	counter = (c + 1)
	mu.Unlock()
	
}
