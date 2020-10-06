package main

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 0
	}
	n := len(dungeon[0])

	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, n), make([]int, n)

	dp[0][n-1] = max(1, 1-dungeon[m-1][n-1])
	for i := n - 2; i >= 0; i-- {
		dp[0][i] = max(1, dp[0][i+1]-dungeon[m-1][i])
	}

	now := 0
	for i := m - 2; i >= 0; i-- {
		now = 1 - now
		dp[now][n-1] = max(1, dp[1-now][n-1]-dungeon[i][n-1])
		for j := n - 2; j >= 0; j-- {
			dp[now][j] = max(1, min(dp[now][j+1], dp[1-now][j])-dungeon[i][j])
		}
	}

	return dp[now][0]
}
