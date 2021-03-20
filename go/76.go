package main

/*
	l-r 当前的子串
	如果满足条件，l++
	如果不满足条件，r++直到满足条件或到末尾
*/

func minWindow(s string, t string) string {
	lens, lent := len(s), len(t)
	if lens < lent {
		return ""
	}

	isNeed := make(map[byte]int, lent)
	m := make(map[byte]int, lent)

	totalNum := 0
	for i := range t {
		if isNeed[t[i]] == 0 {
			totalNum++
		}
		isNeed[t[i]]++
	}

	nowNum := 0
	reslen := len(s)
	res := ""
	l, r := 0, 0
	if isNeed[s[0]] != 0 {
		m[s[0]] = 1
		if m[s[0]] == isNeed[s[0]] {
			nowNum++
		}
	}

	for r < lens {
		if nowNum == totalNum {
			if r-l+1 <= reslen {
				reslen = r - l + 1
				res = s[l : r+1]
			}
			if isNeed[s[l]] != 0 {
				m[s[l]]--
				if m[s[l]] < isNeed[s[l]] {
					nowNum--
				}
			}
			l++
		} else {
			r++
			if r < lens && isNeed[s[r]] != 0 {
				if m[s[r]] == isNeed[s[r]]-1 {
					nowNum++
				}
				m[s[r]]++
			}
		}
	}

	return res
}
