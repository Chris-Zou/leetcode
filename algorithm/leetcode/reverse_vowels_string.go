package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("hello"))
}

func isVowels(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	} else {
		return false
	}

}

func reverseVowels(s string) string {
	l := len(s)
	i, j := 0, l-1
	for i < j {
		for i < j && isVowels(s[i]) == false {
			i++
		}
		for i < j && isVowels(s[j]) == false {
			j--
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
}
