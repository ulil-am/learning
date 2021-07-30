package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	result := make([]int, N)
	max := 0
	for i := range A {
		if A[i] >= 1 && A[i] <= N {
			result[A[i]-1] += 1
			if result[A[i]-1] > max {
				max = result[A[i]-1]
			}
		} else if A[i] == N+1 {
			for i := range result {
				result[i] = max
			}
		}
	}
	return result
}
