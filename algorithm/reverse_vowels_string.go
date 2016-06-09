package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("aA"))
}

func isVowels(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	} else if ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	} else {
		return false
	}

}

func reverseVowels(s string) string {
	l := len(s)
	i, j := 0, l-1

	ret := make([]byte, l)
	for k := 0; k < l; k++ {
		ret[k] = s[k]
	}
	for i < j {
		for i < j && isVowels(s[i]) == false {

			i++
		}
		for i < j && isVowels(s[j]) == false {

			j--
		}
		if i < j {
			ret[i], ret[j] = s[j], s[i]
			i++
			j--
		}
	}
	return string(ret)
}
