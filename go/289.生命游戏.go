/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */
// @lc code=start

// 2:原来是0要下次为1 3:原来是1下次为0的

func isAlive(status int) bool {
	return status == 1 || status == 3
}

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			alive := 0
			if j-1 >= 0 {
				if i-1 >= 0 && isAlive(board[i-1][j-1]) {
					alive++
				}
				if isAlive(board[i][j-1]) {
					alive++
				}
				if i+1 < m && isAlive(board[i+1][j-1]) {
					alive++
				}
			}
			if i-1 >= 0 && isAlive(board[i-1][j]) {
				alive++
			}
			if i+1 < m && isAlive(board[i+1][j]) {
				alive++
			}
			if j+1 < n {
				if i-1 >= 0 && isAlive(board[i-1][j+1]) {
					alive++
				}
				if isAlive(board[i][j+1]) {
					alive++
				}
				if i+1 < m && isAlive(board[i+1][j+1]) {
					alive++
				}
			}
			if board[i][j] == 0 {
				if alive == 3 {
					board[i][j] = 2
				}
			} else {
				if alive < 2 || alive > 3 {
					board[i][j] = 3
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

// @lc code=end

