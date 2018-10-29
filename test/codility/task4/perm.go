package main

import "fmt"

func main() {
	A := []int{3, 1, 2, 4}
	hasil := Solution(A)
	fmt.Printf("%v\n", hasil)
}

func Solution(A []int) int {
	// write your code in Go 1.4
	B := make([]int, len(A))
	for value := range A {
		B[value] = 1
	}
	for j := 1; j <= len(B); j++ {
		if B[j] == 0 {
			return 0
		}
	}
	return 1
}
