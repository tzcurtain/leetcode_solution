package main

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res, now := 0, 1
	tmpM := make(map[int]bool, len(nums))
	for i := range nums {
		tmpM[nums[i]] = true
	}
	var keys []int
	for i := range tmpM {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	lenk := len(keys)
	for i := 0; i < lenk; i++ {
		if i+1 < lenk && keys[i+1] == keys[i]+1 {
			now++
		} else {
			res = max(now, res)
			now = 1
		}
	}
	res = max(now, res)

	return res
}
