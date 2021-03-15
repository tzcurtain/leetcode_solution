/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	var res []int
	appear := make(map[int]int, 10)
	for _, now := range nums {
		appear[now]++
	}
	// 变成求前k大

	return res
}

// @lc code=end

