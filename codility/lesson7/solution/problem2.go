package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	var location, fishSize []int
	for i, val := range A {
		if i == 0 {
			location = append(location, B[i])
			fishSize = append(fishSize, val)
		} else if B[i] != location[len(location)-1] {
			if location[len(location)-1] == 0 {
				location = append(location, B[i])
				fishSize = append(fishSize, val)
			} else {
				if val > fishSize[len(fishSize)-1] {
					fishSize, location = eatFish(fishSize, location)
					location = append(location, B[i])
					fishSize = append(fishSize, val)
				}
			}
		} else {
			location = append(location, B[i])
			fishSize = append(fishSize, val)
		}
		// fmt.Printf("current size and loc %v, %v\n", fishSize, location)
	}

	fishLeft := len(fishSize)

	return fishLeft
}

func eatFish(fS []int, fL []int) ([]int, []int) {
	fSeaten := fS[:len(fS)-1]
	fLoc := fL[:len(fL)-1]

	return fSeaten, fLoc
}
