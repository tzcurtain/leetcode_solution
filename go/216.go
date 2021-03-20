package main

import (
	"sort"
)

/*
	等于是q40数组是1-9的特殊情况
	最后对结果筛选一下就好
*/

func sliceEquals3(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

func combinationSum3(k int, n int) [][]int {
	m := make(map[int][][]int, n+2)

	for j := 1; j <= 9; j++ {
		for i := n; i >= 1; i-- {
			if i == j {
				arr := []int{j}
				m[i] = append(m[i], arr)
			} else if i > j {
				for _, kk := range m[i-j] {
					if len(kk) >= k {
						continue
					}
					kkk := make([]int, len(kk))
					copy(kkk, kk)
					kkk = append(kkk, j)
					sort.Ints(kkk)
					insert := true
					for _, p := range m[i] {
						if sliceEquals3(kkk, p) {
							insert = false
							break
						}
					}
					if insert {
						m[i] = append(m[i], kkk)
					}
				}
			}
			// fmt.Println(i, ":", m[i])
		}
	}

	res := make([][]int, 0)
	for _, val := range m[n] {
		if len(val) == k {
			res = append(res, val)
		}
	}

	return res
}
