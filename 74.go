package main

/*
	两次二分
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	l, r := 0, m-1
	for l < r {
		mid := (l + r) / 2
		if matrix[mid][0] > target {
			r = mid - 1
		} else if matrix[mid][0] == target {
			return true
		} else {
			l = mid + 1
		}
	}

	if l > 0 && matrix[l][0] > target {
		l--
	}
	nowm := l

	l, r = 0, n-1
	for l < r {
		mid := (l + r) / 2
		if matrix[nowm][mid] > target {
			r = mid - 1
		} else if matrix[nowm][mid] == target {
			return true
		} else {
			l = mid + 1
		}
	}

	if matrix[nowm][l] == target {
		return true
	}

	return false
}
