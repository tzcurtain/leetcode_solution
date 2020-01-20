package main

func wordBreak(s string, wordDict []string) []string {
	m := make(map[int]map[string]int, len(s)+1)
	unreachable := make(map[int]bool, len(s))
	wordBreakRecur(s, 0, wordDict, &unreachable, &m)
	var res []string
	for i := range m[len(s)] {
		res = append(res, i)
	}
	return res
}

func wordBreakRecur(s string, l int, wordDict []string, unreachable *map[int]bool, nowDict *(map[int]map[string]int)) bool {
	if l == len(s) {
		return true
	}

	if (*unreachable)[l] {
		return false
	}

	// fmt.Println(l)

	ok := false
	for i := range wordDict {
		nowlen := len(wordDict[i])
		if len(s)-l >= nowlen {
			if s[l:l+nowlen] == wordDict[i] {
				if (*nowDict)[l+nowlen] != nil && (*nowDict)[l+nowlen][wordDict[i]] == 1 {
					continue
				}
				if l == 0 {
					if (*nowDict)[nowlen] == nil {
						(*nowDict)[nowlen] = make(map[string]int, 0)
					}
					(*nowDict)[nowlen][wordDict[i]] = 1
				} else {
					if (*nowDict)[l+nowlen] == nil {
						(*nowDict)[l+nowlen] = make(map[string]int, 0)
					}
					for j := range (*nowDict)[l] {
						tmpStr := j + " " + wordDict[i]
						(*nowDict)[l+nowlen][tmpStr] = 1
					}
				}
				if wordBreakRecur(s, l+nowlen, wordDict, unreachable, nowDict) {
					ok = true
				}
			}
		}
	}

	(*unreachable)[l] = !ok
	return ok
}
