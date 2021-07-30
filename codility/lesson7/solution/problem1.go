package solution

// you can also use imports, for example:
import (
	// "fmt"
	"strings"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	var nested, kriwil, round, block, lengthBlock, lengthRound, lengthKriwil int
	splittedString := strings.Split(S, "")
	var posOK, posOR, posOB []int
	for i := range splittedString {
		switch splittedString[i] {
		case "{":
			posOK = append(posOK, i)
			kriwil += 1
			// fmt.Printf("posOK %v\n", posOK)
		case "[":
			block += 1
			posOB = append(posOB, i)
			// fmt.Printf("posOB %v\n", posOB)
		case "(":
			round += 1
			posOR = append(posOR, i)
			// fmt.Printf("posOR %v\n", posOR)
		case "}":
			if kriwil != 0 {
				lengthBlock = len(posOB)
				lengthRound = len(posOR)
				lengthKriwil = len(posOK)
				if lengthKriwil == 0 {
					return 0
				} else {
					if lengthBlock != 0 && posOB[lengthBlock-1] > posOK[lengthKriwil-1] {
						return 0
					} else if lengthRound != 0 && posOR[lengthRound-1] > posOK[lengthKriwil-1] {

						return 0
					}
					posOK = decreaseSlice(posOK)
					// fmt.Printf("posOK %v\n", posOK)
				}
				kriwil -= 1
			} else {
				return 0
			}
		case "]":
			if block != 0 {
				lengthBlock = len(posOB)
				lengthRound = len(posOR)
				lengthKriwil = len(posOK)
				if lengthBlock == 0 {
					return 0
				} else {
					if lengthKriwil != 0 && posOK[lengthKriwil-1] > posOB[lengthBlock-1] {
						return 0
					} else if lengthRound != 0 && posOR[lengthRound-1] > posOB[lengthBlock-1] {

						return 0
					}
					posOB = decreaseSlice(posOB)
					// fmt.Printf("posOB %v\n", posOB)

				}
				block -= 1
			} else {
				return 0
			}
		case ")":
			if round != 0 {
				lengthBlock = len(posOB)
				lengthRound = len(posOR)
				lengthKriwil = len(posOK)
				if lengthRound == 0 {
					return 0
				} else {
					if lengthKriwil != 0 && posOK[lengthKriwil-1] > posOR[lengthRound-1] {
						return 0
					} else if lengthBlock != 0 && posOB[lengthBlock-1] > posOR[lengthRound-1] {

						return 0
					}
					posOR = decreaseSlice(posOR)
					// fmt.Printf("posOR %v\n", posOR)
				}
				round -= 1
			} else {
				return 0
			}
		}
	}
	nested = 1
	// fmt.Printf("posOR %v\n", posOR)
	if kriwil != 0 || round != 0 || block != 0 {
		nested = 0
	}
	return nested
}

func decreaseSlice(s []int) []int {
	sLen := len(s)
	return s[:sLen-1]
}

// ({{({}[]{})}}[]{})
