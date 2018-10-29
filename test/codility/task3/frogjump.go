package task3

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	var i int
	for i = 0; X < Y; i++ {
		X += D
	}
	return i
}
