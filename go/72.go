package main

/*
	dp[i][j] = dp[i-1][j-1] (word1[i] == word2[j])
			 = min(dp[i-1][j],dp[i-1][j-1],dp[i][j-1]) + 1 (word1[i] != word2[j])
*/

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	if len1*len2 == 0 {
		return len1 + len2
	}
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}
	for i := 1; i < len2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1])
			}
		}
	}

	return dp[len1][len2]
}
