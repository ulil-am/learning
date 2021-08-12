package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
    // init:= false
	var max_ending, max_slice int
	for _, val := range A {
        if len(A) == 1 {
			return val
        } else {
			max_ending = Max(0, max_ending+val)
			max_slice = Max(max_slice, max_ending)
            // init= true
        }
	}
    sort.Ints(A)
    if A[len(A)-1] < 0 {
        max_slice = A[len(A)-1]
    }
	return max_slice
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

[-10]
[-2,1]
[-5,3,2,1]
[-5,-3,-2,-1]
[-2, -2]