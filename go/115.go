package main

// f[i][j] = s[:i]中出现t[:j]的个数
// f[i][j] = f[i-1][j]
// f[i][j] += f[i-1][j-1]  (s[i-1] = t[j-1])

func numDistinct(s string, t string) int {
	lens, lent := len(s), len(t)
	f := make([][]int, lens+1)
	for i := range f {
		f[i] = make([]int, lent+1)
	}

	for i := 0; i <= lens; i++ {
		f[i][0] = 1
	}
	for i := 1; i <= lent; i++ {
		f[0][i] = 0
	}

	for i := 1; i <= lens; i++ {
		for j := 1; j <= lent; j++ {
			f[i][j] = f[i-1][j]
			if s[i-1] == t[j-1] {
				f[i][j] += f[i-1][j-1]
			}
		}
	}

	return f[lens][lent]
}
