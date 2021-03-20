package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := make(map[int]int, 0)
	for i := range nums {
		if t > len(m) {
			for j := range m {
				if m[j] != 0 && abs(j-nums[i]) <= t && (i+1)-m[j] <= k {
					return true
				}
			}
		} else {
			for j := nums[i] - t; j <= nums[i]+t; j++ {
				if m[j] != 0 && (i+1)-m[j] <= k {
					return true
				}
			}
		}
		m[nums[i]] = i + 1
	}
	return false
}
