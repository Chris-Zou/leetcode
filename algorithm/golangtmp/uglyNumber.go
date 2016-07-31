package main

import (
	"fmt"
)

func isUgly(num int) bool {

	for num > 1 && num&1 == 0 {
		num = num >> 1
	}
	for num > 1 && num%3 == 0 {
		num = num / 3
	}
	for num > 1 && num%5 == 0 {
		num = num / 5
	}
	if num == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(2))
	fmt.Println(isUgly(3))
	fmt.Println(isUgly(4))
	fmt.Println(isUgly(5))
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(7))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(9))
	fmt.Println(isUgly(10))
	fmt.Println(isUgly(12))
	fmt.Println(isUgly(13))
	fmt.Println(isUgly(14))
	fmt.Println(isUgly(15))
	fmt.Println(isUgly(16))
	fmt.Println(isUgly(17))
}
