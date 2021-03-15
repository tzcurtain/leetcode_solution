/*
 * @lc app=leetcode.cn id=1503 lang=golang
 *
 * [1503] 所有蚂蚁掉下来前的最后一刻
 */

// @lc code=start
func getLastMoment(n int, left []int, right []int) int {
	res := -1
	for i := 0; i < len(left); i++ {
		if left[i] > res {
			res = left[i]
		}
	}
	for j := 0; j < len(right); j++ {
		if (n - right[j]) > res {
			res = n - right[j]
		}
	}
	return res
}

// @lc code=end

