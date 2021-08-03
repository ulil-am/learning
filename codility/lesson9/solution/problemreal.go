package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	var max1, max3, max2, index1, index2, sumAll int
	var index3 int
	currentmax := 0
	for _, val := range A {
		sumAll += val
	}
	for i, val2 := range A {
		if i == 0 {
			max1 = sumAll - val2
		} else {
			currentmax = sumAll - val2
			if currentmax > max1 {
				max1 = currentmax
				index1 = i
			}
		}
	}
	fmt.Printf("max1, ")
	init3 := false
	for j, val3 := range A {
		if j == index1 {
			continue
		}
		if !(init3) {
			max2 = max1 - val3
			init3 = true
		} else {
			currentmax = max1 - val3
			if currentmax > max2 {
				max2 = currentmax
				index2 = j
			}
		}
	}
	init := false
	for k, val4 := range A {
		if k == index1 || k == index2 {
			continue
		}
		if !(init) {
			init = true
			max3 = max2 - val4
		} else {
			currentmax = max2 - val4
			if currentmax > max3 {
				max3 = currentmax
				index3 = k
			}
		}

	}
	fmt.Printf("index 1,2,3 : %v %v %v", index1, index2, index3)

	return max3

}
