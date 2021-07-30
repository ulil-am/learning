package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "strings"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	var nested, round, lengthRound int
	splittedString := strings.Split(S, "")
	var posOR []int
	for i := range splittedString {
		switch splittedString[i] {
		case "(":
			round += 1
			posOR = append(posOR, i)

		case ")":
			if round != 0 {
				lengthRound = len(posOR)
				if lengthRound == 0 {
					return 0
				}
				posOR = decreaseSlice(posOR)
				round -= 1
			} else {
				return 0
			}
		}
	}
	nested = 1
	if round != 0 {
		nested = 0
	}
	return nested
}

func decreaseSlice(s []int) []int {
	sLen := len(s)
	return s[:sLen-1]
}
