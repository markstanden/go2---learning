// Package summy is a gash math package
package summy

// MySum sums a unlimited number of ints
func MySum(xs ...int) int{
	var sum int
	for _, v := range xs {
		sum += v
	}
	return sum
}