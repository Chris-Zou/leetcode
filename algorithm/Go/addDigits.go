package main

import (
	"fmt"
)

func main() {
	ret := addDigits(1234)
	if ret != 1 {
		fmt.Println("Wrong Answer")
	}
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	ret := 0
	for num != 0 {
		ret = ret + num%10
		num = num / 10
	}

	return addDigits(ret)
}
