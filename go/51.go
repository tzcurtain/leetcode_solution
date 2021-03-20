package main

func solveNQueens(n int) [][]string {
	var res [][]string

	tmp := make([][]byte, n)

	// rowUsed := make([]bool, n)
	colUsed := make([]bool, n)
	diagDaleUsed := make([]bool, 4*n)
	diagHillUsed := make([]bool, 2*n)
	dfsQueen(0, n, &colUsed, &diagHillUsed, &diagDaleUsed, &tmp, &res)

	return res
}

func dfsQueen(now, n int, colUsed, diagHillUsed, diagDaleUsed *[]bool, tmp *[][]byte, res *[][]string) {
	if now == n {
		var nowRes []string
		for _, v := range *tmp {
			nowRes = append(nowRes, string(v))
		}
		*res = append(*res, nowRes)
		return
	}

	if (*tmp)[now] == nil {
		(*tmp)[now] = make([]byte, n)
		for i := 0; i < n; i++ {
			(*tmp)[now][i] = '.'
		}
	}

	for i := 0; i < n; i++ {
		// fmt.Println(now, i, " ", i+now, now-i+2*n)
		if !(*colUsed)[i] && !(*diagHillUsed)[i+now] && !(*diagDaleUsed)[now-i+2*n] {
			(*colUsed)[i], (*diagHillUsed)[i+now], (*diagDaleUsed)[now-i+2*n] = true, true, true
			(*tmp)[now][i] = 'Q'
			dfsQueen(now+1, n, colUsed, diagHillUsed, diagDaleUsed, tmp, res)
			(*tmp)[now][i] = '.'
			(*colUsed)[i], (*diagHillUsed)[i+now], (*diagDaleUsed)[now-i+2*n] = false, false, false
		}
	}
}
