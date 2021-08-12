package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int, C []int) int {
	// write your code in Go 1.4
	allNailed := false
	mapNailed := make(map[int]int)
	var maxNailsPos int
	// var nailed []int
	for i := 0; i < len(C); i++ {
		if i != 0 && maxNailsPos > C[i] {
			continue
		}
		for j := 0; j < len(A); j++ {
			if mapNailed[C[i]] != 0 {
				continue
			}
			if C[i] >= A[j] && C[i] <= B[j] {
				mapNailed[C[i]] += 1
				if j == len(A)-1 {
					allNailed = true
				}
			}
		}
		if allNailed {
			break
		}

		maxNailsPos = C[i]
	}
	// fmt.Printf("mapNailed %v\n", mapNailed)
	min := len(mapNailed)
	if !allNailed {
		min = -1
	}

	return min
}
