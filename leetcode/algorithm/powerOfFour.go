package main

import (
	"fmt"
)

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	if num&(num-1) != 0 {
		return false
	}
	index := 0
	for num != 0 {
		index++
		num = num >> 1
	}

	if index&1 == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(6))
	fmt.Println(isPowerOfFour(7))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(9))
	fmt.Println(isPowerOfFour(10))
	fmt.Println(isPowerOfFour(11))
	fmt.Println(isPowerOfFour(12))
	fmt.Println(isPowerOfFour(13))
	fmt.Println(isPowerOfFour(14))
	fmt.Println(isPowerOfFour(15))
	fmt.Println(isPowerOfFour(16))
}
