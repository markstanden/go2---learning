package main

import (
	"fmt"
	"summy"
	"testing"
)

func TestMySum(t *testing.T){
	testVal := summy.MySum(1,2,3,4,5,6,7,8,9,10)
	if  testVal != 55{
		t.Errorf("mySum(1,2,3,4,5,6,7,8,9,10) expected 55, got %v", testVal)
	}
}

func ExampleMySum() {
	fmt.Println(summy.MySum(1, 2, 3))
	// Output:
	// 6
}
