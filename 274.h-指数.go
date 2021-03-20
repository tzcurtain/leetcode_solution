/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func hIndex(citations []int) int {
	n := len(citations)
	hash := make([]int, n+1)
	for i := 0; i < n; i++ {
		hash[min(n, citations[i])]++
	}
	now := 0
	for i := n; i >= 0; i-- {
		now += hash[i]
		if now >= i {
			return i
		}
	}
	return 0
}

// @lc code=end

