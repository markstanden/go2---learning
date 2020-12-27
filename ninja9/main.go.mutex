package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("Arch:\t\t", runtime.GOARCH)
	fmt.Println("Number of CPUs :", runtime.NumCPU())
	count := 0
	maxGR := 0
	const gs = 1000000

	for i := 0; i < gs; i++ {
		wg.Add(1)
		go incrementor(&count, &maxGR)

	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println("Max Goroutines", maxGR)
}

func incrementor(counter *int, gr *int) {
	runtime.Gosched()
	mu.Lock()

	*counter++
	if runtime.NumGoroutine() > *gr {
		*gr = runtime.NumGoroutine()
	}
	go func() int {
		if runtime.NumGoroutine() > *gr {
			*gr = runtime.NumGoroutine()
		}
		return 20 * 19 * 18 * 17 * 16 * 15 * 14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	}()
	mu.Unlock()
	wg.Done()
}
