package main

import (
	"runtime"
)


func main() {
counter := 0

go func (c int) {

runtime.Gosched()
counter = c + 1

} (counter)

}