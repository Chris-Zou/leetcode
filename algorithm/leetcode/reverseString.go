package main

import (
	"fmt"
)

func reverseString(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	var ret []byte
	for i := l - 1; i >= 0; i-- {
		ret = append(ret, s[i])
	}
	return string(ret)
}

func main() {
	ret := reverseString("hello")
	fmt.Println(ret)
	fmt.Println(reverseString("your"))
}
