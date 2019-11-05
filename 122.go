package main

// 找递增
// func maxProfit(prices []int) int {
// 	if prices == nil || len(prices) == 0 {
// 		return 0
// 	}
// 	res := 0
// 	for i := 0; i < len(prices); i++ {
// 		if i+1 < len(prices) && prices[i+1] > prices[i] {
// 			buy, sell := prices[i], 0
// 			now := i + 1
// 			for now < len(prices) && prices[now] > prices[now-1] {
// 				sell = prices[now]
// 				now++
// 			}
// 			// fmt.Println(sell, buy)
// 			res += sell - buy
// 			i = now - 1
// 		}
// 	}
// 	return res
// }
