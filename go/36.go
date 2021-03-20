package main

func isValidSudoku(board [][]byte) bool {
	columnMap := [10][10]byte{}
	for idx := range board {
		rowMap := [10]byte{}
		for i, val := range board[idx] {
			if val == '.' {
				continue
			}
			val = val - '0'
			if rowMap[val] != 0 || columnMap[i][val] != 0 {
				return false
			}
			rowMap[val] = 1
			columnMap[i][val] = 1
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			totalMap := [10]byte{}
			for p := 3 * i; p < 3*i+3; p++ {
				for q := 3 * j; q < 3*j+3; q++ {
					if board[p][q] == '.' {
						continue
					}
					if totalMap[board[p][q]-'0'] != 0 {
						return false
					}
					totalMap[board[p][q]-'0'] = 1
				}
			}
		}
	}

	return true
}

// func main() {
// 	// return isValidSudoku()
// }
