package main

import (
	"fmt"

	"./dog"
)

func main() {
	var hy float64
	hy = 40 + (10 / 12)
	fmt.Println(dog.Years(hy))
}
