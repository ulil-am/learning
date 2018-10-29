package main

import (
	"fmt"
)

func main() {
	A := []int{9, 3, 9, 3, 9, 7, 9}
	hasil := Solution(A)
	fmt.Printf("%v\n", hasil)
}

// func Solution(A []int, K int) []int {
// 	// write your code in Go 1.4
// 	for i := 1; i <= K; i++ {
// 		for j := len(A) - 1; j > 0; j-- {
// 			if j == len(A)-1 {
// 				A[j], A[0] = A[0], A[j]
// 			} else {
// 				A[j+1], A[j] = A[j], A[j+1]
// 			}
// 		}
// 	}
// 	return A
// }

// func Solution(A []int) int {
// 	// write your code in Go 1.4
// 	length := len(A)
// 	var B = make([]string, length)
// 	var index int
// 	var tag int
// 	for i := range A {
// 		tag = 0
// 		if B[i] == "tagged" {
// 			continue
// 		} else {
// 			for j := i + 1; j < len(A); j++ {
// 				if B[j] == "tagged" {
// 					continue
// 				} else {
// 					if A[i] == A[j] {
// 						B[j] = "tagged"
// 						tag = 1
// 						break
// 					}
// 				}
// 			}
// 		}
// 		if tag == 0 {
// 			index = i
// 			break
// 		}
// 	}
// 	return A[index]
// }

func Solution(A []int) int {
	// write your code in Go 1.4
	length := len(A)
	var B = make([]string, length)
	var index int
	var tag int
	for i := range A {
		tag = 0
		if B[i] == "tagged" {
			continue
		} else {
			for j := i + 1; j < len(A); j++ {
				if B[j] == "tagged" {
					continue
				} else {
					if A[i] == A[j] {
						B[j] = "tagged"
						tag = 1
						break
					}
				}
			}
		}
		if tag == 0 {
			index = i
			break
		}
	}
	return A[index]
}
