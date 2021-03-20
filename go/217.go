package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, 0)
	for i := range nums {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
