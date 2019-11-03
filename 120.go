package main

func minimumTotal(triangle [][]int) int {
	lenf := len(triangle)
	f := make([]int, lenf)
	res := IntMax

	for i := range f {
		f[i] = IntMax
	}
	f[0] = 0

	for i := 0; i < lenf; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				f[0] += triangle[i][0]
			} else {
				f[j] = min(f[j], f[j-1]) + triangle[i][j]
			}
			if i == lenf-1 {
				if f[j] < res {
					res = f[j]
				}
			}
		}
	}

	return res
}
