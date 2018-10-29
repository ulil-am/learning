package main

import "fmt"

func main() {
	A := 11
	B := 345
	K := 17
	hasil := Solution(A, B, K)
	fmt.Printf("%v\n", hasil)
}

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	// count := 0
	// for i := A; i <= B; i++ {
	// 	if i%K == 0 {
	// 		count++
	// 	}
	// }
	count := (B - A) / K
	if A%K == 0 || B%K == 0 {
		count += 1
	}
	return count
}
