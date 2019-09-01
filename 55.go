package main
/*
	q45的简化版
*/
func canJump(nums []int) bool {
	lenn := len(nums)
	now := nums[0]
	for i := 1; i < lenn; i++ {
		if i+nums[i] <= now {
			continue
		}
		if i <= now && i+nums[i] > now {
			now = i + nums[i]
		} else {
			return false
		}
	}
	return true
}
