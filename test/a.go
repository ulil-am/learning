package main

import "fmt"
import "strings"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	s := string{"photo.jpg, Warsaw, 2013-09-05 14:08:15
		john.png, London, 2015-06-20 15:13:22
		myFriends.png, Warsaw, 2013-09-05 14:07:13
		Eiffel.jpg, Paris, 2015-07-23 08:03:02
		pisatower.jpg, Paris, 2015-07-22 23:59:59
		BOB.jpg, London, 2015-08-05 00:02:03
		notredame.png, Paris, 2015-09-01 12:00:00
		me.jpg, Warsaw, 2013-09-06 15:40:22
		a.png, Warsaw, 2016-02-13 13:33:50
		b.jpg, Warsaw, 2016-01-02 15:12:22
		c.jpg, Warsaw, 2016-01-02 14:34:30
		d.jpg, Warsaw, 2016-01-02 15:15:01
		e.png, Warsaw, 2016-01-02 09:49:09
		f.png, Warsaw, 2016-01-02 10:55:32
		g.jpg, Warsaw, 2016-02-29 22:13:11"}
	fmt.Printf("%s\n", Solution(s))
}

func Solution(S string) string {
	// write your code in Go 1.4
	var stringhasil string
	
	// var filename, city, date []string
	splitted:= strings.Split(S,"\n")
	for i :=range splitted {
		splitted2 := strings.Split(splitted, ",")
		filename[i], city[i], date[i] := splitted2[0], splitted2[1], splitted2[2]
		for j := range filename {
			splitFile := filename[j].Split(splitted, ".")
			file[j],ext[j] := splitFile[0], splitFile[1]
		}
	for indexcity := range city {

	}
		
	}

	return stringhasil
}

// func Solution(A []int) int {
//     // write your code in Go 1.4
//     minHasil := 1
//     var B []int
//     B = make([]int, 1000001)
//     // var hasil int
//     // // sorted := Sort(A)
//     if len(A) == 1 && A[0] !=1 {
//         return 1
//     }
//     for i, value := range A {
//         if A[i]>0 {
//             B[value]++
//         }
//     }
//     for j:=1;j<len(B);j++ {
//         if B[j]==0 {
//             minHasil= j
//             break
//         }
//     }
//     if minHasil < 1 {
//         minHasil = 1
//     }
//     return minHasil
// }


func Sort(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	panjang := len(A) / 2
	left := Sort(A[:panjang])
	right := Sort(A[panjang:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	sortb := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(sortb, right...)
		}
		if len(right) == 0 {
			return append(sortb, left...)
		}
		if left[0] <= right[0] {
			sortb = append(sortb, left[0])
			left = left[1:]
		} else {
			sortb = append(sortb, right[0])
			right = right[1:]
		}
	}

	return sortb
}
