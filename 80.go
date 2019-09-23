package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	now := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[now] && nums[i] == nums[now-1] {
			continue
		} else {
			now++
			nums[now] = nums[i]
		}
	}

	return now + 1
}
