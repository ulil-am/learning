package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.package solution

// you can also use imports, for example:
// import "fmt"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sortA := A
	sort.Ints(sortA)
	length := len(A)
	max := 0
	var finalmax int
	var index1, index2, times3 int
	maxLen := sortA[length-1] * sortA[length-2]
	minLen := sortA[0] * sortA[1]
	// maxMin := sortA[0] * sortA[length-1]
	loop := 1

	if sortA[length-1] > 0 && sortA[0] < 0 && sortA[1] < 1 && sortA[length-3] <= 0 {
		finalmax = sortA[length-1] * sortA[0] * sortA[1]
	} else if sortA[length-1] < 0 || sortA[0] > 0 {
		finalmax = sortA[length-1] * sortA[length-2] * sortA[length-3]
	} else {
		if maxLen > minLen {
			max = maxLen
			index1 = length - 1
			index2 = length - 2
		} else if minLen >= maxLen {
			max = minLen
			index1 = 0
			index2 = 1
		}
		for j, val := range sortA {
			if j != index1 && j != index2 {

				times3 = max * val
				if loop == 1 || times3 > finalmax {
					finalmax = times3
				}
				loop++
			}
		}
	}

	return finalmax
}
