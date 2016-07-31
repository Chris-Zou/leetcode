package main

import (
	"fmt"
)

func main() {
	fmt.Println("1534236469", reverse(1534236469), " -123", reverse(-123))
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
		ret = -ret
		if ret < -1<<31 {
			return 0
		}
		return ret
	} else {
		if ret > 1<<31-1 {
			return 0
		}
		return ret
	}
}
