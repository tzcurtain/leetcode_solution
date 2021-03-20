package main

func rangeBitwiseAnd(m int, n int) int {
	var bitM, bitN []int
	for m > 0 {
		bitM = append(bitM, m%2)
		m /= 2
	}
	for n > 0 {
		bitN = append(bitN, n%2)
		n /= 2
	}
	lenM, lenN := len(bitM), len(bitN)
	if lenM != lenN {
		return 0
	}

	nowi, nowj := lenM-1, lenN-1
	res := 0
	for nowi >= 0 && nowj >= 0 && bitM[nowi] == bitN[nowj] {
		nowi--
		nowj--
	}
	base := 1
	for i := 0; i <= nowi; i++ {
		base *= 2
	}
	for i := nowi + 1; i < lenM; i++ {
		res += base * bitM[i]
		base *= 2
	}
	return res
}
