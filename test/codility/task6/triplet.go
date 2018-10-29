package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var maks2 int
	var maksAll int
	var indeks1, indeks2 int
	for i := range A {
		for j := i + 1; j < len(A); j++ {
			if i == 0 {
				maks2 = A[i] * A[j]
				indeks1 = i
				indeks2 = j
			} else {
				if A[i]*A[j] > maks2 {
					maks2 = A[i] * A[j]
					indeks1 = i
					indeks2 = j
				}
			}
		}
	}
	nowmaks := maks2
	for k := 0; k < len(A); k++ {
		if k == indeks1 || k == indeks2 {
			continue
		}
		maksAll = A[k] * maks2
		if maksAll > nowmaks {
			nowmaks = maksAll
		}
	}
	return nowmaks
}
