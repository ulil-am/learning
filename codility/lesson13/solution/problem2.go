package solution

// you can also use imports, for example:
import (
	"fmt"
	"math"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) []int {
	// write your code in Go 1.4
	var result []int
	var position, mod, pow int
	for i, val := range A {
		position = calcFibonacci(val)
		pow = int(math.Pow(2, float64(B[i])))
		// fmt.Printf("pow %v\n", pow)
		mod = modulus(position, pow)
		result = append(result, mod)
	}
	// fmt.Printf("result %v\n", result)
	return result
}

func calcFibonacci(N int) int {
	// // fmt.Printf("N %v\n", N)
	// fibonacci := make([]int, N+2)
	// fibonacci[0] = 0
	// fibonacci[1] = 1
	// for i := 2; i <= N+1; i++ {
	// 	fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	// 	// fmt.Printf("fibonacci %v\n", fibonacci)
	// }
	// return fibonacci[N+1]
	fibonacci := (math.Pow((1+math.Sqrt(float64(5)))/2, float64(N+1)) - math.Pow((1-math.Sqrt(float64(5)))/2, float64(N+1))) / math.Sqrt(float64(5))
	// fib, frac := math.Modf(fibonacci)
	// if frac < 1 && frac > 0 {
	// 	fibonacci = fib + 1
	// }
	fibonacci = Round(fibonacci)
	fmt.Printf("fibonacci %v\n ", fibonacci)

	return int(fibonacci)
}

func modulus(x int, m int) int {
	r := x % m
	if r < 0 {
		r = r + m
	}
	return r
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func RoundToEven(x float64) float64 {
	t := math.Trunc(x)
	odd := math.Remainder(t, 2) != 0
	if d := math.Abs(x - t); d > 0.5 || (d == 0.5 && odd) {
		return t + math.Copysign(1, x)
	}
	return t
}

// ([5,6,7,8,9,10,11,12,13,14,15], [3, 2, 4, 3, 1,3, 2, 4, 3, 1 ,1])
