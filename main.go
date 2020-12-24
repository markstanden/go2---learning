package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

		wg.Add(1)
	go func () {
		fmt.Println("Get your hands off of my woman")
		wg.Done()
	} ()
wg.Add(1)
	go func() {
		
		fmt.Println("Motherfucker")
		wg.Done()
	} ()

	wg.Wait()
}