package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	var fishleft, fishPos []int
	var totPos int
	for i, val := range A {
		if i == 0 {
			fishleft = append(fishleft, val)
			fishPos = append(fishPos, B[i])
		} else {
			totPos = len(fishPos)
			// totLeft = len(fishleft)
			if B[i] != fishPos[totPos-1] && B[i] == 1 {
				fishleft = append(fishleft, val)
				fishPos = append(fishPos, B[i])
			} else if B[i] != fishPos[totPos-1] && B[i] == 0 {
				fishleft, fishPos = eatFish(val, B[i], fishleft, fishPos)
			} else {
				fishleft = append(fishleft, val)
				fishPos = append(fishPos, B[i])
			}
		}
		fmt.Printf("=======main fishleft, fish post %v, %v\n", fishleft, fishPos)
	}
	return len(fishleft)
	// write your code in Go 1.4
}

func eatFish(A int, B int, fishleft []int, fishPos []int) (left []int, pos []int) {
	aSurvive := false
	lenFish1 := len(fishleft) - 1
	if fishleft[lenFish1] < A {
		fishleft[lenFish1] = A
		fishPos[lenFish1] = B
		aSurvive = true
	}
	fmt.Printf("fishleft, fish post %v, %v %v\n", fishleft, fishPos, aSurvive)
	var i int
	if aSurvive && len(fishPos) > 1 {
		for i = (len(fishleft) - 2); i >= 0; {
			fmt.Printf("fishleft, fish post %v, %v %v\n", fishleft, fishPos, aSurvive)
			if fishPos[i] == 1 && fishleft[i] < A {
				fishleft[i] = A
				fishPos[i] = B
				left = fishleft[0 : i+1]
				pos = fishPos[0 : i+1]
				i--
				fmt.Printf("after if fishleft, fish post %v, %v %v\n", fishleft, fishPos, aSurvive)
			} else if fishPos[i] == 1 && fishleft[i] > A {
				aSurvive = false
				left = fishleft[0:i]
				pos = fishPos[0:i]
				// i--
				fmt.Printf("after else if1 fishleft, fish post %v, %v %v\n", fishleft, fishPos, aSurvive)
				break
			} else if fishPos[i] == 0 {
				left = fishleft[0 : i+2]
				pos = fishPos[0 : i+2]
				fmt.Printf("after elseif2 fishleft, fish post %v, %v %v\n", fishleft, fishPos, aSurvive)
				break
			}
			fmt.Printf("after fishleft, fish post %v, %v %v\n", fishleft, fishPos, aSurvive)
		}
	} else {
		left = fishleft
		pos = fishPos
	}
	// left = fishleft[0 : i+1]
	// pos = fishPos[0 : i+1]
	// x := len(left)

	return
}
