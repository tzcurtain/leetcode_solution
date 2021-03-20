/*
 * @lc app=leetcode.cn id=1492 lang=golang
 *
 * [1492] n 的第 k 个因子
 */

// @lc code=start
func kthFactor(n int, k int) int {
	nowk := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			nowk++
			if nowk == k {
				return i
			}
		}
	}

	if nowk == k-1 {
		return n
	}
	return -1
}

// @lc code=end

