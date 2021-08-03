package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
// import "sort"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	// sortedA := A
	mapValue := make(map[int]int)
	index := -1
	maxOccurence := 0
	// maxVal := 0
	// sort.Ints(sortedA)
	for i, val := range A {
		mapValue[val] += 1
		if mapValue[val] > maxOccurence {
			maxOccurence = mapValue[val]
			index = i
		}
	}
	// fmt.Printf("maxOccurence %v\n", maxOccurence)

	minLeader := len(A)/2 + 1

	if maxOccurence < minLeader {
		// fmt.Printf("masuk sini")
		index = -1
	}
	return index
}
