package main

func minCut(s string) int {
	lens := len(s)
	dp := make([][]bool, lens)
	for i := range dp {
		dp[i] = make([]bool, lens)
	}

	m := make(map[int]int, 1)
	for k := 1; k <= lens; k++ {
		for i := 0; i <= lens-k; i++ {
			j := i + k - 1
			dp[i][j] = s[i] == s[j] && (k < 3 || dp[i+1][j-1])
		}
	}
	return minCutHelper(s, 0, dp, m)

}

func minCutHelper(s string, start int, dp [][]bool, m map[int]int) int {
	if _, ok := m[start]; ok {
		return m[start]
	}

	lens := len(s)
	if dp[start][lens-1] {
		return 0
	}
	mind := IntMax
	for i := start; i < lens; i++ {
		if dp[start][i] {
			mind = min(mind, 1+minCutHelper(s, i+1, dp, m))
		}
	}
	m[start] = mind
	return mind
}
