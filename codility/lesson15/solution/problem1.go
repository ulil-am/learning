package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	mapA := make(map[int]int)
	var absVal int
	for _, val := range A {
		absVal = int(math.Abs(float64(val)))
		if mapA[absVal] == 0 {
			mapA[absVal]++
		}
	}
	return len(mapA)
}
