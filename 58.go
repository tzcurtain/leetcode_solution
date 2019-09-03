package main

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	res := 0
	for i >= 0 && !isAlpha(s[i]) {
		i--
	}
	for i >= 0 && isAlpha(s[i]) {
		i--
		res++
	}

	return res
}
