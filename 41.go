package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := range nums {
		ni := abs(nums[i]) - 1
		if ni < n {
			nums[ni] = -abs(nums[ni])
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
