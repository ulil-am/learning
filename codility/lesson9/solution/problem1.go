package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) []int {
	// write your code in Go 1.4
	var result []int
	var nonDivisor int
	for i, val := range A {
		nonDivisor = 0
		for j, val2 := range A {
			if j == i {
				continue
			}
			if val < val2 || (val > val2 && val%val2 != 0) {
				nonDivisor += 1
			}
		}
		result = append(result, nonDivisor)
	}
	return result
}
