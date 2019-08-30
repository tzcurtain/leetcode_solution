package main

func myPow(x float64, n int) float64 {
	switch {
	case n < 0:
		return 1 / myPow(x, -n)
	case n == 0:
		return 1
	case n == 1:
		return x
	case n&1 == 0:
		res := myPow(x, n>>1)
		return res * res
	default: // n%2 == 0
		res := myPow(x, n>>1)
		return res * res * x
	}
}
