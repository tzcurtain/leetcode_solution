/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	f := make([]int, amount+2)

	for i := 0; i < len(coins); i++ {
		if coins[i] > amount {
			continue
		}
		f[coins[i]] = 1
	}

	for j := 0; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			if j < coins[i] || f[j-coins[i]] == 0 {
				continue
			}
			if f[j] == 0 {
				f[j] = f[j-coins[i]] + 1
			} else {
				f[j] = min(f[j], f[j-coins[i]]+1)
			}
		}
	}
	if f[amount] == 0 {
		f[amount] = -1
	}

	return f[amount]
}

// @lc code=end

