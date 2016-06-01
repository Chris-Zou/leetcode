package main

import (
	"fmt"
)

func getSingleOne(n int) int {
	ret := 0
	for n != 0 {
		if n%10 == 1 {
			ret++
		}
		n = n / 10
	}

	return ret
}

func getOnes(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		ret = ret + getSingleOne(i)
	}

	return ret
}

func countDigitOne(n int) int {

	if n < 0 {
		return 0
	}

	ret := 0
	base, b := 1, 1
	for n > 0 {
		tmp := 0
		if n%10 == 1 {
			tmp = 1
		}
		ret = ret + ((n + 8) / 10 * base) + tmp*b
		b = b + n%10*base
		base *= 10
		n = n / 10
	}

	return ret
}

func check(n int) bool {
	ret := getOnes(n)
	tmp := countDigitOne(n)
	if ret == tmp {
		return true
	} else {
		return false
	}
}

func main() {
	/*fmt.Println(getOnes(10))
	fmt.Println(getOnes(11))
	fmt.Println(getOnes(12))
	fmt.Println(getOnes(13))*/

	fmt.Println(getOnes(555))
	fmt.Println(getOnes(100))

	fmt.Println(check(100))
	fmt.Println(check(555))
	fmt.Println(check(101))
	fmt.Println(check(99))
	fmt.Println(check(10))
	fmt.Println(check(567))
	fmt.Println(check(987))
	fmt.Println(check(654))
	fmt.Println(check(110))
	fmt.Println(check(111))
}
