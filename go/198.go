package main

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var f [2][2]int // 0 偷 1 不偷

	f[0][0], f[0][1] = nums[0], 0
	now := 0

	for i := 1; i < n; i++ {
		now = 1 - now
		f[now][0] = f[1-now][1] + nums[i]
		f[now][1] = max(f[1-now][0], f[1-now][1])
	}

	return max(f[now][0], f[now][1])
}
