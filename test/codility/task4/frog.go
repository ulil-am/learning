package main

import (
	"fmt"
)

func main() {
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	N := 5
	hasil := Solution(N, A)
	fmt.Printf("%v\n", hasil)
}

type pos struct {
	Position int
	index    int
}

type ByPosition []pos

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	if X > 100000 {
		return -1
	}

	isi := make([]pos, len(A))
	minCount := -1
	maxIndex := -1
	// farest :=0
	// var B [][]int
	B := make([]int, X+1)

	for i := 0; i < len(A); i++ {
		isi[i].Position = A[i]
		isi[i].index = i
		B[A[i]]++
	}

	for l := 1; l < len(B)-1; l++ {
		if B[l+1]-B[l] > 1 {
			return -1
		}
	}
	sorted := (ByPosition(isi))

	for j := 0; j < len(sorted); j++ {
		// if sorted[j+1].Position-sorted[j].Position > 1 {
		// 	return -1
		// }
		if sorted[j].index > maxIndex {
			maxIndex = sorted[j].index
		}
		if sorted[j].Position == X {
			if minCount == -1 || sorted[j].index < maxIndex {
				minCount = maxIndex
			} else if sorted[j].index > maxIndex {
				minCount = sorted[j].index
				break
			}
		}
	}
	return minCount
}
