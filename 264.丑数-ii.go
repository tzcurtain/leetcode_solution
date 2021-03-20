/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func nthUglyNumber(n int) int {
	// 2 3 2*2 5 2*3 2*2*2 2*5 2*2*3 3*5 2*2*2*2 2*3*3 2*2*5 5*5 2*3*5
	arr := []int{1}
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		now := min(arr[i2]*2, min(arr[i3]*3, arr[i5]*5))
		arr = append(arr, now)
		if now == arr[i2]*2 {
			i2++
		}
		if now == arr[i3]*3 {
			i3++
		}
		if now == arr[i5]*5 {
			i5++
		}
	}
	return arr[n-1]
}

// @lc code=end

