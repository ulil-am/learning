package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var peakIndex []int
	// var currentFlags int
	for i, val := range A {
		if i == 0 || i == len(A)-1 {
			continue
		} else if val > A[i-1] && val > A[i+1] {
			fmt.Printf("masuk sini val %v\n", val)
			peakIndex = append(peakIndex, i)
		}
	}
	k := len(peakIndex)
	if k == 0 {
		return 0
	}
	fmt.Printf("peakIndex %v\n", peakIndex)
	return k
}
