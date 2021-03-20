package main

func isHappy(n int) bool {
	appear := make(map[int]bool, 1)

	for n != 1 && !appear[n] {
		appear[n] = true
		res := 0
		for n > 0 {
			res += (n % 10) * (n % 10)
			n /= 10
		}
		n = res
	}

	return n == 1
}
