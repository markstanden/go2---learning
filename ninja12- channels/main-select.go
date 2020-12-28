package main

import "fmt"

//create channel c as global int channel
var e = make(chan num)
var o = make(chan num)

type num int

func (n num) send(even, odd chan<- num) {
	if n%2 == 0 {
		even <- n
	} else { 
		odd <- n
	}
}

func receive(even, odd <-chan num) {
	select {
	case v := <-even:
		fmt.Println("Even:\t", v)
		
	case v := <-odd:
		fmt.Println("Odd:\t", v)
	}

}

func main() {
	total := 100
	for i := 0; i < total; i++ {
		go num(i).send(e, o)
	}

	for i := 0; i < total; i++ {
		receive(e, o)
	}


}
