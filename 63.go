package main

/*
	just like 62
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, m)

	for i := range f {
		f[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		} else {
			f[i][0] = 1
		}
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		} else {
			f[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}
