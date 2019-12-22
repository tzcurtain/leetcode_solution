package main

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'Q'
			}
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'Q' {
			dfsBoard(board, i, 0, m, n)
		}
		if board[i][n-1] == 'Q' {
			dfsBoard(board, i, n-1, m, n)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'Q' {
			dfsBoard(board, 0, j, m, n)
		}
		if board[m-1][j] == 'Q' {
			dfsBoard(board, m-1, j, m, n)
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if board[i][j] == 'Q' {
				board[i][j] = 'X'
			}
		}
	}

}

func dfsBoard(board [][]byte, x, y, m, n int) {
	board[x][y] = 'O'
	if x-1 > 0 && board[x-1][y] == 'Q' {
		dfsBoard(board, x-1, y, m, n)
	}
	if x+1 < m && board[x+1][y] == 'Q' {
		dfsBoard(board, x+1, y, m, n)
	}
	if y-1 > 0 && board[x][y-1] == 'Q' {
		dfsBoard(board, x, y-1, m, n)
	}
	if y+1 < n && board[x][y+1] == 'Q' {
		dfsBoard(board, x, y+1, m, n)
	}
}
