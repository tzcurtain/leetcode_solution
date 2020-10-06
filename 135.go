package main

import "sort"

func candy(ratings []int) int {
	m := make(map[int][]int, 0)
	total := 0

	lenR := len(ratings)

	for i := 0; i < lenR; i++ {
		if m[ratings[i]] == nil {
			m[ratings[i]] = make([]int, 1)
			m[ratings[i]][0] = i
		} else {
			m[ratings[i]] = append(m[ratings[i]], i)
		}
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	tmp := make([]int, lenR)

	for i := range keys {
		nowArr := m[i]
		for j := range nowArr {
			nowPos := nowArr[j]
			shouldPut := 1
			if nowPos-1 >= 0 && ratings[nowPos] > ratings[nowPos-1] {
				shouldPut = tmp[nowPos-1] + 1
			}
			if nowPos+1 < lenR && ratings[nowPos] > ratings[nowPos+1] {
				shouldPut = max(shouldPut, tmp[nowPos+1]+1)
			}
			total += shouldPut
			tmp[nowPos] = shouldPut
		}
	}

	return total
}
