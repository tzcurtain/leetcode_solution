package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for i := range nums {
		if m[nums[i]] != 0 && (i+1)-m[nums[i]] <= k {
			return true
		}
		m[nums[i]] = i + 1
	}
	return false
}
