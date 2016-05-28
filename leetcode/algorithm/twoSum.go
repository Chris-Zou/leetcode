package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		k, ok := mp[target-nums[i]]
		if ok && i != k {
			return []int{i, k}
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}
