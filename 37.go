package main

var usedRowVal, usedColumnVal, usedDistrictVal, fixed [12][12]bool
var solved = false

/*
	回溯解决
*/

// 全局变量导致测试和实际提交结果不同
func solveSudoku(board [][]byte) {
	usedRowVal, usedColumnVal, usedDistrictVal, fixed = [12][12]bool{}, [12][12]bool{}, [12][12]bool{}, [12][12]bool{}
	solved = false
	initConfig(board)
	if board[0][0] > '0' && board[0][0] <= '9' {
		i := board[0][0] - '0'
		fixed[0][0] = false
		usedRowVal[0][i], usedColumnVal[0][i], usedDistrictVal[0][i] = false, false, false
		tryToSolveSudoku(board, 0, 0)
	} else {
		for i := 1; i <= 9; i++ {
			if !usedRowVal[0][i] && !usedColumnVal[0][i] && !usedDistrictVal[0][i] {
				board[0][0] = byte(i) + '0'
				tryToSolveSudoku(board, 0, 0)
				if solved {
					break
				}
				board[0][0] = '.'
				usedRowVal[0][i], usedColumnVal[0][i], usedDistrictVal[0][i] = false, false, false
			}
		}
	}
}
func initConfig(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] > '0' && board[i][j] <= '9' {
				fixed[i][j] = true
				val := board[i][j] - '0'
				district := i/3*3 + j/3
				usedRowVal[i][val], usedColumnVal[j][val], usedDistrictVal[district][val] = true, true, true
			}
		}
	}
}
func tryToSolveSudoku(board [][]byte, x, y int) {
	if solved || x == 9 || (x == 8 && y == 8) {
		solved = true
		return
	}
	val := board[x][y] - '0'
	nowDistrict := x/3*3 + y/3
	usedRowVal[x][val], usedColumnVal[y][val], usedDistrictVal[nowDistrict][val] = true, true, true
	nextx, nexty := x, y+1
	if nexty == 9 {
		nextx, nexty = nextx+1, 0
	}
	for fixed[nextx][nexty] {
		nexty = nexty + 1
		if nexty == 9 {
			nextx, nexty = nextx+1, 0
		}
	}
	if nextx == 9 {
		solved = true
		return
	}
	nextDistrict := nextx/3*3 + nexty/3
	for i := 1; i <= 9; i++ {
		if !usedRowVal[nextx][i] && !usedColumnVal[nexty][i] && !usedDistrictVal[nextDistrict][i] {
			board[nextx][nexty] = byte(i) + '0'
			tryToSolveSudoku(board, nextx, nexty)
			if solved {
				break
			}
			board[nextx][nexty] = '.'
			usedRowVal[nextx][i], usedColumnVal[nexty][i], usedDistrictVal[nextDistrict][i] = false, false, false
		}
	}
}

// func main() {

// 	// d := [][]byte{
// 	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// 	// }

// 	d := [][]byte{{'.', '.', '9', '7', '4', '8', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '2', '.', '1', '.', '9', '.', '.', '.'}, {'.', '.', '7', '.', '.', '.', '2', '4', '.'}, {'.', '6', '4', '.', '1', '.', '5', '9', '.'}, {'.', '9', '8', '.', '.', '.', '3', '.', '.'}, {'.', '.', '.', '8', '.', '3', '.', '2', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '6'}, {'.', '.', '.', '2', '7', '5', '9', '.', '.'}}

// 	solveSudoku(d)

// 	fmt.Println("-----res-----")
// 	for i := range d {
// 		for j := range d[i] {
// 			if d[i][j] > '0' && d[i][j] <= '9' {
// 				fmt.Print(d[i][j] - '0')
// 			} else {
// 				fmt.Print("0")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
