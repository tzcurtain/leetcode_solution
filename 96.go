package main

/*
	f[0] = 1
	f[n] = f[n-i] * f[i-1] (1 <= i <= n)
*/

func numTrees(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	for now := 1; now <= n; now++ {
		f[now] = 0
		for i := 1; i <= now; i++ {
			f[now] += f[now-i] * f[i-1]
		}
	}
	return f[n]
}
