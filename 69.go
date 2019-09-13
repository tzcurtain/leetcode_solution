package main

func mySqrt(x int) int {
	for i := 0; ; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}
}
