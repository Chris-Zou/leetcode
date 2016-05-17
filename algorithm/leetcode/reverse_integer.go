package main

import (
	"fmt"
)

func main() {
	fmt.Println("123", reverse(123), " -123", reverse(-123))
}

func reverse(x int) int {
	flag := true
	if x < 0 {
		x = -x
		flag = false
	}
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x = x / 10
	}

	if flag == false {
		return -ret
	} else {
		return ret
	}
}
