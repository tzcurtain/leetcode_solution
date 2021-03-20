/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		res ^= nums[i]
	}
	for i := 1; i <= n; i++ {
		res ^= i
	}
	return res
}

// @lc code=end

