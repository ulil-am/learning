package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(K int, A []int) int {
	// write your code in Go 1.4
	var totalrope, curTotal, curSum, j int
	for i := 0; i < len(A); {
		// fmt.Printf("loop i ke %v\n", i)
		curTotal = 0
		if A[i] >= K {
			totalrope += 1
			i++
			continue
		}

		for j = i + 1; j < len(A); j++ {
			// fmt.Printf("loop j ke %v\n", j)
			if j == i+1 {
				curSum = A[i] + A[j]
			} else {
				curSum += A[j]
			}

			if curSum >= K {
				curTotal += 1
				break
			}
		}
		i += j + 1
		totalrope += curTotal
	}

	return totalrope
}
