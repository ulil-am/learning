package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(H []int) int {
	// write your code in Go 1.4
	var blocks, newBlock int
	// height := make(map[int]int)
	var index []int
	for i, val := range H {
		if i == 0 {
			index = append(index, val)
		} else {
			if val > H[i-1] {
				index = append(index, val)
			} else if val < H[i-1] {
				index, newBlock = checkBlock(index, val)
				blocks += newBlock
			}
		}
		// fmt.Printf("loop ke- %v, index, blocks  %v, %v \n", i, index, blocks)
	}
	if len(index) != 0 {
		blocks += len(index)
	}
	return blocks
}

func checkBlock(index []int, h int) (blockLeft []int, closedBlock int) {
	var i int
	for i = (len(index) - 1); i >= 0; {
		// fmt.Printf("checkBLock ke- %v, index, closedBlock  %v, %v \n", i, index, closedBlock)
		if index[i] > h {
			closedBlock += 1
			i--
		} else if index[i] == h {
			i--
		} else {
			break
		}
	}
	// fmt.Printf("i %v \n", i)
	if i > -1 {
		blockLeft = index[0 : i+1]
	}
	blockLeft = append(blockLeft, h)
	// fmt.Printf("blockLeft and closedBlock %v, %v \n", blockLeft, closedBlock)
	return
}
