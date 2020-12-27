package main

import "testing"

func TestMySum(t *testing.T){
	testVal := mySum(1,2,3,4,5,6,7,8,9,10)
	if  testVal != 55{
		t.Errorf("mySum(1,2,3,4,5,6,7,8,9,10) expected 55, got %v", testVal)
	}
}