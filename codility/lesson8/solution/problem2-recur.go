package solution

// you can also use imports, for example:
// import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var equiLeaders int
	// var leader int
	equiLeaders = getLeader(A, 0)

	// for i := range A {
	// 	if i == len(A)-1 {
	// 		continue
	// 	}
	// 	fmt.Printf("Divide slice %v, %v\n", A[0:i+1], A[i+1:])
	// 	leader1, leader1exist := getLeader(A[0 : i+1])

	// 	leader2, leader2exist := getLeader(A[i+1:])
	// 	fmt.Printf("leader %v, %v\n", leader1, leader2)
	// 	if !leader1exist || !leader2exist {
	// 		continue
	// 	}
	// 	if leader1 == leader2 {
	// 		equiLeaders++
	// 	}
	// }
	return equiLeaders

}

func getLeader(A []int, l int) int {
	var equiLeaders int
	aLen := len(A)

	if l != aLen-2 {
		equiLeaders = getLeader(A, l+1)
	} else {
		return 0
	}

	B := (A[0 : l+1])
	C := (A[l+1:])
	hasLeader := true
	mapValue := make(map[int]int)
	mapValue2 := make(map[int]int)
	var maxNumber, maxNumber2 int
	maxOccurence := 0
	for _, val := range B {
		mapValue[val] += 1
		if mapValue[val] > maxOccurence {
			maxOccurence = mapValue[val]
			maxNumber = val
		}
	}
	minLeader := len(B)/2 + 1
	if maxOccurence < minLeader {
		// fmt.Printf("masuk sini")
		maxNumber = -1
		hasLeader = false
	}

	maxOccurence = 0
	for _, val2 := range C {
		mapValue2[val2] += 1
		if mapValue2[val2] > maxOccurence {
			maxOccurence = mapValue2[val2]
			maxNumber2 = val2
		}
	}
	// fmt.Printf("maxOccurence %v\n", maxOccurence)

	minLeader = len(C)/2 + 1
	hasLeader2 := true
	if maxOccurence < minLeader {
		// fmt.Printf("masuk sini")
		maxNumber2 = -1
		hasLeader2 = false
	}

	if maxNumber == maxNumber2 && hasLeader2 && hasLeader {
		equiLeaders++
	}
	return equiLeaders
}

//[1, 2, 3, 4, 5]
