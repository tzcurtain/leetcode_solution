package main

// 1-n有x个质数
// x! * (n-x)! mod 10^9 + 7

func numPrimeArrangements(n int) int {
	if n == 1 {
		return 1
	}
	primeNum := 1 // 2
	for i := 3; i <= n; i++ {
		if isPrime(i) {
			primeNum++
		}
	}

	res := 1
	for i := 2; i <= primeNum; i++ {
		res = res * i % Modint
	}
	for i := 2; i <= n-primeNum; i++ {
		res = res * i % Modint
	}

	return res
}
