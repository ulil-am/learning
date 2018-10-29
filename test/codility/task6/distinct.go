package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sorted := MergeSort(A)
	counter := 0
	var initial int
	for k := 0; k < len(sorted); k++ {
		if k == 0 {
			initial = sorted[k]
			counter++
		} else {
			if sorted[k] != initial {
				initial = sorted[k]
				counter++
			}
		}
	}
	return counter
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

func main() {

	A := []int{2, 1, 1, 2, 3, 1}
	// N := 5
	hasil := Solution(A)
	fmt.Printf("%v\n", hasil)
}
