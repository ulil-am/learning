package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, M int) int {
	// write your code in Go 1.4
	chocoEaten := 0
	finish := false
	currentChoco := 0
	eated := make(map[int]bool)
	for finish != true {
		if eated[currentChoco] == false {
			eated[currentChoco] = true
			chocoEaten += 1
			currentChoco += M
			if currentChoco >= N {
				currentChoco = currentChoco % N
			}
		} else {
			finish = true
		}
	}
	return chocoEaten
}
