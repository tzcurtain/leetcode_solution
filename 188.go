package main

// f[i][j] 第i天进行j次操作能获取的利润 j奇数为买，偶数为卖
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}
	f := make([][]int, 2)

	totallen := min(2*k+1, n+1)

	f[0], f[1] = make([]int, totallen), make([]int, totallen)

	f[0][0], f[0][1] = 0, -prices[0]
	now := 0
	nowk := 1
	for i := 1; i < n; i++ {
		now = 1 - now
		f[now][0] = 0
		f[now][1] = max(f[1-now][1], -prices[i])
		if nowk < totallen-1 {
			nowk++
		}
		for j := 2; j < nowk; j++ {
			if j%2 == 0 { // 卖出
				f[now][j] = max(f[1-now][j], f[1-now][j-1]+prices[i])
			} else {
				f[now][j] = max(f[1-now][j], f[1-now][j-1]-prices[i])
			}
		}
		if nowk%2 == 0 {
			if nowk < i+1 {
				f[now][nowk] = max(f[1-now][nowk], f[1-now][nowk-1]+prices[i])
			} else {
				f[now][nowk] = f[1-now][nowk-1] + prices[i]
			}
		} else {
			if nowk < i+1 {
				f[now][nowk] = max(f[1-now][nowk], f[1-now][nowk-1]-prices[i])
			} else {
				f[now][nowk] = f[1-now][nowk-1] - prices[i]
			}
		}
	}

	res := 0
	for i := 0; i <= totallen-1; i++ {
		res = max(res, f[now][i])
	}

	return res
}
