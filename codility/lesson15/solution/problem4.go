package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	// lenA := len(A)
	var sumFront, sumBack int
	var min, back int
	for i := range A {
		back = len(A) - 1
		for j := i; j <= len(A)/2; j++ {
			sumFront = A[i] + A[j]
			sumBack = A[i] + A[back]
			if j == 0 {
				// sum = A[i]+A[j]
				min = int(math.Abs(float64(sumFront)))
				if min > int(math.Abs(float64(sumBack))) {
					min = int(math.Abs(float64(sumBack)))
				}
			} else if min > int(math.Abs(float64(sumBack))) {
				min = int(math.Abs(float64(sumBack)))
			} else if min > int(math.Abs(float64(sumFront))) {
				min = int(math.Abs(float64(sumFront)))
			}
			back -= 1
		}
	}
	return min
}
