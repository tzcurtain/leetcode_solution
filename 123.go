package main

/*
	dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
	dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
*/

// func maxProfit(prices []int) int {
// 	if prices == nil || len(prices) == 0 {
// 		return 0
// 	}
// 	n := len(prices)
// 	f0, f1, f2, f3 := -prices[0], IntMin, IntMin, IntMin

// 	for i := 1; i < n; i++ {
// 		// fmt.Println(f0, f1, f2, f3)
// 		f0 = max(f0, -prices[i])
// 		f1 = max(f1, f0+prices[i])
// 		f2 = max(f2, f1-prices[i])
// 		f3 = max(f3, f2+prices[i])
// 	}

// 	return max(0, f3)
// }
