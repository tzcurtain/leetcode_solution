package main

import "fmt"

type pair struct {
	a, b int
}

func isMatch(s string, p string) bool {
	m := make(map[pair]bool)
	return isStringMatch(s, p, 0, 0, m)
}

/* 类似第10题 */

func isStringMatch(s, p string, i, j int, m map[pair]bool) bool {
	fmt.Println("i: ", i, " j:", j)
	val, ok := m[pair{i, j}]

	if ok == true {
		return val
	}
	ans := false
	if i == len(s) || j >= len(p) {
		for j < len(p) && p[j] == '*' {
			j++
		}
		return i == len(s) && j == len(p)
	}
	if p[j] == '?' {
		ans = isStringMatch(s, p, i+1, j+1, m)
	} else if p[j] == '*' {
		ans = isStringMatch(s, p, i+1, j, m) || isStringMatch(s, p, i, j+1, m)
	} else {
		ans = s[i] == p[j] && isStringMatch(s, p, i+1, j+1, m)
	}

	m[pair{i, j}] = ans

	return ans
}
