/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	pre1, pre2 := 0, slow
	for pre1 != pre2 {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}
	return pre1
}

// @lc code=end

