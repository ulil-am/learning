package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(M int, A []int) int {
	// write your code in Go 1.4
	var mapCheck map[int]int
	var total int
	// notDisc := false
	for i := range A {
		mapCheck = make(map[int]int)
		for j := i; j < len(A); j++ {
			if mapCheck[A[j]] == 0 {
				mapCheck[A[j]]++
				total += 1
			} else {
				break
			}
		}
	}
	return total
}

// package solution

// // you can also use imports, for example:
// // import "fmt"
// // import "os"

// // you can write to stdout for debugging purposes, e.g.
// // fmt.Println("this is a debug message")

// func Solution(M int, A []int) int {
// 	// write your code in Go 1.4
// 	var mapCheck map[int]int
// 	var total, breakIndex, totalFinal int
//     totalI := make([]int, len(A))
// 	// notDisc := false
// 	for i := range A {
//         if i!=0 && i<breakIndex {
//             totalI[i]=totalI[i-1]-1
//         } else {
//             mapCheck = make(map[int]int)
//             total = 0
//             for j := i; j < len(A); j++ {
//                 if mapCheck[A[j]] == 0 {
//                     mapCheck[A[j]]++
//                     total += 1
//                 } else {
//                     breakIndex = j
//                     break
//                 }
//             }
//         }

//         totalI = append(totalI, total)
// 	}
//     for _, val2 := range totalI {
//         totalFinal+=val2
//     }
// 	return totalFinal
// }
