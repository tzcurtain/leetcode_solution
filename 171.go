package main

func titleToNumber(s string) int {
	res := 0
	base := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[i]-'A') * base
		base *= 26
	}

	return res
}
