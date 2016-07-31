package main

import (
	"fmt"
)

func countBits(num int) []int {
	var ret []int
	ret = append(ret, 0)
	for i := 1; i <= num; i++ {
		ret = append(ret, ret[i&(i-1)]+1)
	}
	return ret
}

func main() {
	ret := countBits(5)
	fmt.Println(ret)

	ret = countBits(123456)
	fmt.Println("finished, length of ret is", len(ret))
}
