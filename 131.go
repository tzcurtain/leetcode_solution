package main

func partition(s string) [][]string {
	lens := len(s)
	if lens == 1 {
		return [][]string{{s}}
	}
	var res [][]string
	for i := 0; i < lens-1; i++ {
		if isPartition(s[:i+1]) {
			tmp := partition(s[i+1:])
			for k := range tmp {
				res = append(res, append([]string{s[:i+1]}, tmp[k]...))
			}
		}
	}

	if isPartition(s) {
		res = append(res, []string{s})
	}

	return res
}

func isPartition(s string) bool {
	lens := len(s)
	for i := 0; lens-1-i > i; i++ {
		if s[i] != s[lens-1-i] {
			return false
		}
	}

	return true
}
