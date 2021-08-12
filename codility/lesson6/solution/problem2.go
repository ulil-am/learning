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
	var finalmax, finalmax1, finalmax2, times3 int
	maxLen := sortA[length-1] * sortA[length-2]
	minLen := sortA[0] * sortA[1]
	// maxMin := sortA[0] * sortA[length-1]
	loop := 1

	if sortA[length-1] > 0 && sortA[0] < 0 && sortA[1] < 1 && sortA[length-3] <= 0 {
		finalmax = sortA[length-1] * sortA[0] * sortA[1]
	} else if sortA[length-1] < 0 || sortA[0] > 0 {
		finalmax = sortA[length-1] * sortA[length-2] * sortA[length-3]
	} else {

		for j, val := range sortA {
			if j != length-1 && j != length-2 {

				times3 = maxLen * val
				if loop == 1 || times3 > finalmax {
					finalmax1 = times3
				}
				loop++
			}
		}

		for j, val := range sortA {
			if j != 0 && j != 1 {

				times3 = minLen * val
				if loop == 1 || times3 > finalmax {
					finalmax2 = times3
				}
				loop++
			}
		}
		finalmax = Max(finalmax1, finalmax2)

	}

	return finalmax
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
