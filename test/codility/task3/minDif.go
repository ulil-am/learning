package main

import "fmt"

func main() {
	A := []int{3, 1, 2, 4, 3}
	hasil := Solution3(A)
	fmt.Printf("%v\n", hasil)
}

func Solution3(A []int) int {
	// write your code in Go 1.4
	// var minDif int
	// // var k int
	// // var j int
	// var dif int
	// var leftSum int
	// var rightSum int
	dif := Calc(A)
	// for i := 1; i < len(A); i++ {
	// 	leftSum = Sum(A[:i])
	// 	rightSum = Sum(A[i:])
	// 	// for j = 0; j < i; j++ {
	// 	// 	leftSum += A[j]
	// 	// }
	// 	// for k = i; k < len(A); k++ {
	// 	// 	rightSum += A[k]
	// 	// }
	// 	dif = Abs(leftSum - rightSum)
	// 	if i == 1 {
	// 		minDif = dif
	// 	} else {
	// 		if dif < minDif {
	// 			minDif = dif
	// 		}
	// 	}
	// 	fmt.Printf("Hasil terkecil P = %v dengan ls = %v dan rs = %v adalah %v\n", i, leftSum, rightSum, dif)
	// }

	return dif
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Sum(isi []int) int {
	if len(isi) <= 1 {
		return isi[0]
	}
	var l int
	var r int
	n := len(isi) / 2
	l += Sum(isi[:n])
	r += Sum(isi[n:])
	return Hitung(l, r)
}

func Hitung(l int, r int) int {
	return l + r
}

func Calc(A []int) int {
	var hasil int
	hasil = 0
	P := 1

	hasil = Sum2(A, P, 0)
	return hasil
}

func Sum2(A []int, P int, minDif int) int {
	var dif int
	leftSum := 0
	rightSum := 0
	if P == len(A) {
		return minDif
	}
	for j := 0; j < P; j++ {
		leftSum += A[j]
	}
	for k := P; k < len(A); k++ {
		rightSum += A[k]
	}
	dif = Abs(leftSum - rightSum)
	if P == 1 {
		minDif = dif
	} else {
		if dif < minDif {
			minDif = dif
		}
	}

	minDif = Sum2(A, P+1, minDif)

	return minDif
}
