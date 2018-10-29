package task3

func Solution2(A []int) int {
	// write your code in Go 1.4

	if len(A) == 0 {
		return 1
	}
	sorted := sort(A)
	var i int
	var pos int
	for i = 0; i < len(A); i++ {
		pos = i + 1
		if sorted[0] != 1 {
			return 1
		} else if sorted[i] != pos {
			return pos
		}
	}
	return pos
}

func sort(A []int) []int {
	for i := 0; i < len(A); i++ {
		minimum := i
		for j := i + 1; j < len(A); j++ {
			if A[minimum] > A[j] {
				minimum = j
			}
		}
		if minimum != i {
			A[i], A[minimum] = A[minimum], A[i]
		}
	}
	return A
}
