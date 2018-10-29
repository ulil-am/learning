package main

import "fmt"

func main() {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	N := 5
	hasil := Solution(N, A)
	fmt.Printf("%v\n", hasil)
}

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	var max int
	var B []int
	max = 0

	// var C []int
	B = make([]int, N)
	for i := 0; i < len(A); i++ {
		if A[i] <= N && A[i] >= 1 {
			B[A[i]-1]++
			if max < B[A[i]-1] {
				max = B[A[i]-1]
			}
		} else if A[i] > N {
			for k := range B {
				B[k] = max
			}
		}
		fmt.Printf("Iterasi ke %v = %v\n", i, B)
	}
	return B
}
