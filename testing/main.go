package main

import (
	"fmt"
)

func main(){
	fmt.Println(mySum(1,2,3,4,5,6,7,8,9,10))
}

func mySum(xs ...int) int{
	var sum int
	for _, v := range xs {
		sum += v
	}
	return sum
}