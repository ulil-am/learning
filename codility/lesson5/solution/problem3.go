package solution

// you can also use imports, for example:
import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    var minPos int
    average := 0.00
    minAverage := 0.00
    var curSum int
    var divider int
    for i:=0; i <len(A)-1;i++ {
        curSum = 0
        for j:=i+1; j <len(A);j++ {
            divider = j-i+1
            // average = 0
            if j == i+1 {
                curSum = A[i]+A[j]
                average = float64(curSum)/float64(divider)
            } else {
                curSum += A[j]
                average = float64(curSum)/float64(divider)
            }
            fmt.Println("average for start pos", i,"  end pos ",j, "are ",average," with current min ",minAverage)
            if (i==0 && j==1) || (average < minAverage) {
                minAverage = average
                minPos = i
            }
        }
    }

    return minPos
}
