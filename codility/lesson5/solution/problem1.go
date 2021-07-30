package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	total := 0
	var mod int

	mod = B % K
	BMod := B - mod
	if BMod >= A {
		if mod == 0 {
			total = ((B - A) / K) + 1
		} else {
			total = ((BMod - A) / K) + 1
		}
	}

	return total
}
