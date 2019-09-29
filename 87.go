package main

/*
	递归
*/

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if !isOfSameCompose(s1, s2) {
		return false
	}

	// fmt.Println(len(s1))
	lens := len(s1)
	if lens <= 2 || s1 == s2 {
		return true
	}

	for i := 0; i < lens-1; i++ {
		// 前缀一样
		if isScramble(s1[:i+1], s2[:i+1]) && isScramble(s1[i+1:lens], s2[i+1:lens]) {
			return true
		}
		// 后缀和前缀一致
		if isScramble(s1[:i+1], s2[lens-i-1:lens]) && isScramble(s1[i+1:lens], s2[:lens-i-1]) {
			return true
		}
	}

	return false
}

func isOfSameCompose(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// fmt.Println(s1, s2)
	m := make(map[byte]int, 0)

	for i := 0; i < len(s1); i++ {
		m[s1[i]] = m[s1[i]] + 1
		m[s2[i]] = m[s2[i]] - 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
