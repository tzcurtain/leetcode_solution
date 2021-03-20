/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */

// @lc code=start
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func hIndex(citations []int) int {
	// h[0] - h[i] - h[n-1]
	// (n-i)篇论文引用次数至少h[i]篇
	n := len(citations)
	if n == 0 {
		return 0
	}
	l, r := 1, n
	for l < r {
		if l == r-1 {
			if citations[n-r] >= r {
				return min(r, citations[n-1])
			}
			return min(l, citations[n-1])
		}
		mid := (l + r) / 2
		// n-mid n
		if citations[n-mid] >= mid {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return min(l, citations[n-1])
}

// @lc code=end

