package main

/*
	q54稍改一下
*/
func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for idx := range matrix {
		matrix[idx] = make([]int, n)
	}

	used := make([][]bool, len(matrix))
	for idx := range used {
		used[idx] = make([]bool, len(matrix[0]))
	}
	moveInSpiral2(&matrix, &used, 0, 0, len(matrix), len(matrix[0]), SpiralRight, 1)
	return matrix
}

func moveInSpiral2(matrix *[][]int, used *[][]bool, l, r, m, n, direction, now int) {
	(*matrix)[l][r] = now
	(*used)[l][r] = true
	switch direction {
	case SpiralRight:
		r++
	case SpiralDown:
		l++
	case SpiralLeft:
		r--
	case SpiralUp:
		l--
	}

	if outOfSpiralBoundary2(l, r, m, n, used) {
		switch direction {
		case SpiralRight:
			r--
			l++
		case SpiralDown:
			l--
			r--
		case SpiralLeft:
			l--
			r++
		case SpiralUp:
			l++
			r++
		}
		if outOfSpiralBoundary2(l, r, m, n, used) {
			return
		}
		direction = (direction + 1) % 4
	}

	moveInSpiral2(matrix, used, l, r, m, n, direction, now+1)
}

func outOfSpiralBoundary2(l, r, m, n int, used *[][]bool) bool {
	return l < 0 || l >= m || r < 0 || r >= n || (*used)[l][r]
}
