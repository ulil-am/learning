package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var total int
	for i, val := range A {
		for j := i + 1; j < len(A)-1; j++ {
			for k := j + 1; k < len(A); k++ {
				fmt.Printf("slice : %v, %v, %v\n", i, j, k)
				if val+A[j] > A[k] && val+A[k] > A[j] && A[k]+A[j] > val {
					total += 1
				}
			}
		}
	}
	return total
}
