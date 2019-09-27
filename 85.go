package main

// 对矩阵每行当成q84来做
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([]int, n)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[j]++
			} else {
				dp[j] = 0
			}
		}
		res = max(res, largestRectangleArea(dp))
		// fmt.Println(res)
	}

	return res
}
