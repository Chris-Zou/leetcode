package main

import (
	"fmt"
)

func getBits(n int) int {
	ret := 0
	for n != 0 {
		n = n & (n - 1)
		ret++
	}
	return ret
}

func countBits(num int) []int {
	var ret []int
	for i := 0; i <= num; i++ {
		ret = append(ret, getBits(i))
	}
	return ret
}

func main() {
	ret := countBits(5)
	fmt.Println(ret)

	ret = countBits(123456)
	fmt.Println("finished, length of ret is", len(ret))
}
