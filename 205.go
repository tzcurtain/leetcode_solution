package main

func isIsomorphic(s string, t string) bool {
	dict := make(map[byte]byte, 0)
	used := make(map[byte]bool, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := dict[s[i]]; ok {
			if dict[s[i]] != t[i] {
				return false
			}
		} else {
			if used[t[i]] {
				return false
			}
			dict[s[i]] = t[i]
			used[t[i]] = true
		}
	}

	return true
}
