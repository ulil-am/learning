package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var equiLeaders int
	for i := range A {
		if i == len(A)-1 {
			continue
		}
		fmt.Printf("Divide slice %v, %v\n", A[0:i+1], A[i+1:])
		leader1, leader1exist := getLeader(A[0 : i+1])

		leader2, leader2exist := getLeader(A[i+1:])
		fmt.Printf("leader %v, %v\n", leader1, leader2)
		if !leader1exist || !leader2exist {
			continue
		}
		if leader1 == leader2 {
			equiLeaders++
		}
	}
	return equiLeaders

}

func getLeader(A []int) (int, bool) {
	hasLeader := true
	mapValue := make(map[int]int)
	var maxNumber int
	maxOccurence := 0
	for _, val := range A {
		mapValue[val] += 1
		if mapValue[val] > maxOccurence {
			maxOccurence = mapValue[val]
			maxNumber = val
		}
	}
	// fmt.Printf("maxOccurence %v\n", maxOccurence)

	minLeader := len(A)/2 + 1

	if maxOccurence < minLeader {
		// fmt.Printf("masuk sini")
		maxNumber = -1
		hasLeader = false
	}
	return maxNumber, hasLeader

}

//[1, 2, 3, 4, 5]
