package main

func calculate(s string) int {
	res, now := 0, 0
	for i := range s {
		if isDigit(s[i]) {
			now = now*10 + (s[i] - '0')
		} else if isSymbol(s[i]) {

		} else if s[i] == ' ' {

		}
	}

	return res
}
