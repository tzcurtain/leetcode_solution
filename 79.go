package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				used[i][j] = true
				if wordDfs(i, j, 1, word, &used, board) {
					return true
				}
				used[i][j] = false
			}
		}
	}

	return false
}

func wordDfs(i, j, now int, word string, used *[][]bool, board [][]byte) bool {
	if now == len(word) {
		return true
	}
	if i-1 >= 0 && !(*used)[i-1][j] && board[i-1][j] == word[now] {
		(*used)[i-1][j] = true
		if wordDfs(i-1, j, now+1, word, used, board) {
			return true
		}
		(*used)[i-1][j] = false
	}
	if i+1 < len(board) && !(*used)[i+1][j] && board[i+1][j] == word[now] {
		(*used)[i+1][j] = true
		if wordDfs(i+1, j, now+1, word, used, board) {
			return true
		}
		(*used)[i+1][j] = false
	}
	if j-1 >= 0 && !(*used)[i][j-1] && board[i][j-1] == word[now] {
		(*used)[i][j-1] = true
		if wordDfs(i, j-1, now+1, word, used, board) {
			return true
		}
		(*used)[i][j-1] = false
	}
	if j+1 < len(board[0]) && !(*used)[i][j+1] && board[i][j+1] == word[now] {
		(*used)[i][j+1] = true
		if wordDfs(i, j+1, now+1, word, used, board) {
			return true
		}
		(*used)[i][j+1] = false
	}

	return false
}
