package main

func rotate(matrix [][]int) {
	l, r := 0, len(matrix[0])-1
	for l < r {
		// [l,l] - [l,r] - [r,r] - [r,l] 移动r-l步
		// swap a & b
		for i := l + 1; i < r; i++ {
			matrix[l][i], matrix[i][r] = matrix[i][r], matrix[l][i]
		}
		// swap a & c
		for i := l + 1; i < r; i++ {
			matrix[l][i], matrix[r][r-(i-l)] = matrix[r][r-(i-l)], matrix[l][i]
		}
		// swap a & d
		for i := l + 1; i < r; i++ {
			matrix[l][i], matrix[r-(i-l)][l] = matrix[r-(i-l)][l], matrix[l][i]
		}
		// swap four edges
		matrix[l][l], matrix[l][r], matrix[r][l], matrix[r][r] =
			matrix[r][l], matrix[l][l], matrix[r][r], matrix[l][r]
		l++
		r--
	}
}
