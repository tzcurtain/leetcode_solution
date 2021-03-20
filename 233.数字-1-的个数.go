/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countDigitOne(n int) int {
	base := 1
	res := 0
	for n >= base {
		res += n/(10*base)*base + min(base, max(n%(10*base)-base+1, 0))
		base *= 10
	}
	return res
}

// @lc code=end

