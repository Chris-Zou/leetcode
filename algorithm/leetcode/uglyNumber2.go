package main

import (
	"fmt"
	"math"
)

var uglyNum []int

func init() {
	uglyNum = append(uglyNum, 1)
	len := 1
	i, j, k := 0, 0, 0
	for uglyNum[len-1] < math.MaxInt32 {
		tmpMin := min(uglyNum[i]*2, uglyNum[j]*3, uglyNum[k]*5)
		uglyNum = append(uglyNum, tmpMin)
		len++

		if uglyNum[i]*2 == tmpMin {
			i++
		}
		if uglyNum[j]*3 == tmpMin {
			j++
		}
		if uglyNum[k]*5 == tmpMin {
			k++
		}
	}
}

func min(a, b, c int) int {
	tmp := a
	if b < tmp {
		tmp = b
	}
	if c < tmp {
		tmp = c
	}
}

func nthUglyNumber(n int) int {
	if n < 1 {
		return 1
	}
	return uglyNum[n-1]
}

func main() {
	for i := 0; i < 11; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")

	fmt.Println(nthUglyNumber(11))
	fmt.Println(nthUglyNumber(12))
	fmt.Println(nthUglyNumber(13))
	fmt.Println(nthUglyNumber(14))
}
