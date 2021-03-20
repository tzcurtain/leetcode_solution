package main

func hammingWeight(num uint32) int {
	total := 0
	if num < 0 {
		if num == uint32(^(^uint32(0) >> 1)) {
			return 32
		}
		num = -num
	}
	for num > 0 {
		if num%2 == 1 {
			total++
		}
		num = num / 2
	}
	return total
}
