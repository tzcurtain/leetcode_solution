package main

func plusOne(digits []int) []int {
	res := make([]int, len(digits))
	now := len(digits) - 1
	digits[now]++
	for now > 0 && digits[now] >= 10 {
		digits[now] = 0
		digits[now-1]++
		now--
	}
	copy(res, digits)
	if res[0] == 10 {
		res[0] = 0
		res = append(res[:0], append([]int{1}, res[0:]...)...)
	}
	return res
}
