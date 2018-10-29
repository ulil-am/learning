package main

// func Solution(A []int) int {
// 	// write your code in Go 1.4
// 	var B []string
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
// 					}
// 				}
// 			}
// 		}
// 		if tag == 1 {
// 			index = i
// 		}
// 	}
// 	return index
// }
