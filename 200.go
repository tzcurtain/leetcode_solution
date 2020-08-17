package main

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}

	res := 0
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				res++
				dfsIsland(grid, i, j, n, m, visited)
			}
		}
	}
	return res
}

func dfsIsland(grid [][]byte, x, y, n, m int, visited [][]bool) {
	visited[x][y] = true
	if x+1 < n && !visited[x+1][y] && grid[x+1][y] == '1' {
		dfsIsland(grid, x+1, y, n, m, visited)
	}
	if x-1 >= 0 && !visited[x-1][y] && grid[x-1][y] == '1' {
		dfsIsland(grid, x-1, y, n, m, visited)
	}
	if y+1 < m && !visited[x][y+1] && grid[x][y+1] == '1' {
		dfsIsland(grid, x, y+1, n, m, visited)
	}
	if y-1 >= 0 && !visited[x][y-1] && grid[x][y-1] == '1' {
		dfsIsland(grid, x, y-1, n, m, visited)
	}
}
