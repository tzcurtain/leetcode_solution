package main

/* 迷宫方向常量 */
const (
	SpiralRight = iota 
	SpiralDown
	SpiralLeft
	SpiralUp
)

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	used := make([][]bool, len(matrix))
	for idx := range used {
		used[idx] = make([]bool, len(matrix[0]))
	}
	moveInSpiral(matrix, &used, 0, 0, len(matrix), len(matrix[0]), SpiralRight, &res)
	return res
}

func moveInSpiral(matrix [][]int, used *[][]bool, l, r, m, n, direction int, res *[]int) {
	// fmt.Println(l, r, m, n)
	*res = append(*res, matrix[l][r])
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

	if outOfSpiralBoundary(l, r, m, n, used) {
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
		if outOfSpiralBoundary(l, r, m, n, used) {
			return
		}
		direction = (direction + 1) % 4
	}

	moveInSpiral(matrix, used, l, r, m, n, direction, res)
}

func outOfSpiralBoundary(l, r, m, n int, used *[][]bool) bool {
	return l < 0 || l >= m || r < 0 || r >= n || (*used)[l][r]
}
