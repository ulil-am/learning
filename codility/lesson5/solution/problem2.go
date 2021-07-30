package solution

// you can also use imports, for example:
// import "fmt"
import "strings"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, P []int, Q []int) []int {
	// write your code in Go 1.4
	M := len(P)
	minImpact := 0
	curImpact := 0
	var result []int
	char := strings.Split(S, "")
	for i := 0; i < M; i++ {
		minImpact = 0
		for j := P[i]; j <= Q[i]; j++ {
			switch char[j] {
			case "A":
				curImpact = 1
			case "C":
				curImpact = 2
			case "G":
				curImpact = 3
			case "T":
				curImpact = 4
			}
			if minImpact == 0 || curImpact < minImpact {
				minImpact = curImpact
			}
		}
		result = append(result, minImpact)
	}
	return result
}
