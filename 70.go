package main

func climbStairs(n int) int {
	f0, f1, f2 := 0, 1, 1
	for i := 2; i <= n; i++ {
		f2 = f1 + f0
		f1 = f2
		f0 = f1
	}
	return f2
}
