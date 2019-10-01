package main

/*
	找规律-镜面反射
*/
func grayCode(n int) []int {
	totalLen := 1
	for i := 1; i <= n; i++ {
		totalLen *= 2
	}
	res := make([]int, totalLen)
	res[0] = 0
	if n == 0 {
		return res
	}
	res[1] = 1

	now := 2
	times := 1
	for times < n {
		for i := 0; i < now; i++ {
			res[now+i] = res[now-i-1] + now
		}
		times++
		now *= 2
	}

	return res
}
