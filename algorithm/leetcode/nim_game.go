package main

import (
	"fmt"
)

func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(canWinNim(1))
	fmt.Println(canWinNim(2))
	fmt.Println(canWinNim(3))
	fmt.Println(canWinNim(4))
	fmt.Println(canWinNim(5))
	fmt.Println(canWinNim(6))
	fmt.Println(canWinNim(7))
}
