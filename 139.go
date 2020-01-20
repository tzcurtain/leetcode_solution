package main

// try memoization
func wordBreak(s string, wordDict []string) bool {
	m := make(map[int]bool, len(s))
	return wordBreakRecur(s, 0, wordDict, &m)
}

func wordBreakRecur(s string, l int, wordDict []string, unreachable *map[int]bool) bool {
	// fmt.Println(l)
	if l == len(s) {
		return true
	}
	if (*unreachable)[l] == true {
		return false
	}
	for i := range wordDict {
		nowlen := len(wordDict[i])
		if len(s)-l >= nowlen {
			if s[l:l+nowlen] == wordDict[i] && wordBreakRecur(s, l+nowlen, wordDict, unreachable) {
				return true
			}
		}

	}
	(*unreachable)[l] = true // l to len(s)-1 is unreachable
	return false
}
