package main

import "fmt"

func main() {
	A := []int{1, 0, 0, 1, 1}
	hasil := Solution(A)
	fmt.Printf("%v\n", hasil)
}

func Solution(A []int) int {
	pair := 0

	total1 := 0
	for i := range A {
		if A[i] == 1 {
			total1++
		}
	}
	for i := range A {
		if A[i] == 0 {
			pair += total1
		} else {
			total1--
		}
	}
	if pair > 1000000000 {
		return -1
	}
	return pair
}
