package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = 1

	for i := 1; i < l; i++ {
		flag := false
		for j := i - dp[i-1]; j < i; j++ {
			if s[i] == s[j] {
				flag = true
				dp[i] = i - j
				break
			}
		}
		if flag == false {
			dp[i] = dp[i-1] + 1
		}
	}
	ret := dp[0]
	for i := 1; i < l; i++ {
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret

}

func main() {
	ret := lengthOfLongestSubstring("pwwkew")
	fmt.Println(ret)
}
