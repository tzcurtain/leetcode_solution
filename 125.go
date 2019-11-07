package main

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if !isAlDigit(s[l]) {
			l++
			continue
		}
		if !isAlDigit(s[r]) {
			r--
			continue
		}
		// fmt.Println(s[l], s[r])
		if s[l] != s[r] {
			if (isUpperCase(s[r]) && isLowerCase(s[l]) && s[l] == s[r]+32) ||
				(isUpperCase(s[l]) && isLowerCase(s[r]) && s[r] == s[l]+32) {
			} else {
				return false
			}
		}
		l++
		r--
	}
	return true
}
