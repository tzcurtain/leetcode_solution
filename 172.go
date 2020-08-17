package main

func trailingZeroes(n int) int {
	fives := 0
	power5 := 5
	for n >= power5 {
		fives += n / power5
		power5 *= 5
	}
	return fives
}
