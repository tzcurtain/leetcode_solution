package main

/*
	Slight change of q51
*/

func totalNQueens(n int) int {
	res := 0
	// rowUsed := make([]bool, n)
	colUsed := make([]bool, n)
	diagDaleUsed := make([]bool, 4*n)
	diagHillUsed := make([]bool, 2*n)
	dfsQueen2(0, n, &colUsed, &diagHillUsed, &diagDaleUsed, &res)
	return res
}

func dfsQueen2(now, n int, colUsed, diagHillUsed, diagDaleUsed *[]bool, res *int) {
	if now == n {
		*res++
		return
	}

	for i := 0; i < n; i++ {
		// fmt.Println(now, i, " ", i+now, now-i+2*n)
		if !(*colUsed)[i] && !(*diagHillUsed)[i+now] && !(*diagDaleUsed)[now-i+2*n] {
			(*colUsed)[i], (*diagHillUsed)[i+now], (*diagDaleUsed)[now-i+2*n] = true, true, true
			dfsQueen2(now+1, n, colUsed, diagHillUsed, diagDaleUsed, res)
			(*colUsed)[i], (*diagHillUsed)[i+now], (*diagDaleUsed)[now-i+2*n] = false, false, false
		}
	}
}
