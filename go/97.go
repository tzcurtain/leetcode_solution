package main

/*
	f[i][j] 表示s1的前i个字符和s2的前j个字符能否组成s3的i+j前缀
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				f[i][j] = true
			} else if i == 0 {
				f[i][j] = f[i][j-1] && (s2[j-1] == s3[i+j-1])
			} else if j == 0 {
				f[i][j] = f[i-1][j] && (s1[i-1] == s3[i+j-1])
			} else {
				f[i][j] = (f[i][j-1] && (s2[j-1] == s3[i+j-1])) ||
					(f[i-1][j] && (s1[i-1] == s3[i+j-1]))
			}
		}
	}

	return f[m][n]
}
