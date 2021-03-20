package main

import "regexp"
/*
	题目很vague导致得多试几次才知道正确的表达式
*/
func isNumber(s string) bool {
    pat := "\\s*[+-]?((\\d+\\.?\\d*)|(\\d*\\.?\\d+))(e[+-]?\\d+)?\\s*"
	re := regexp.MustCompile(pat)
	match := re.Find([]byte(s))
	return string(match) == s
}
