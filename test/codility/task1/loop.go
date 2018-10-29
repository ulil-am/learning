package main

import (
	"fmt"
	"strconv"
)

func main() {
	N := 1041
	hasil := Solution(N)
	fmt.Printf("%v\n", hasil)
}

func Solution(N int) int {
	// write your code in Go 1.4
	n := int64(N)
	var binary string
	binary = strconv.FormatInt(n, 2)
	fmt.Printf("Hasil convert ke binary = %v\n", binary)
	gap := 0
	sequence := 0
	maxSequence := 0
	for i := range binary {
		// bytebin := int64(binary[i])
		// binary2 := strconv.FormatInt(bytebin, 2)
		cek := string(binary[i])
		fmt.Printf("Binary = %v, i = %v, yang dicek = %s\n", binary, i, cek)
		if cek == "0" {
			sequence++
		} else if cek == "1" && i != 0 {
			gap++
			fmt.Printf("Gap ke = %v\n", gap)
			if sequence > maxSequence {
				maxSequence = sequence
				fmt.Printf("Max gap ke %v = %v\n", gap, maxSequence)
			}
			sequence = 0
		}
	}
	if gap == 0 {
		return 0
	}
	fmt.Printf("Hasil akhir = %v\n", maxSequence)
	return maxSequence
}
