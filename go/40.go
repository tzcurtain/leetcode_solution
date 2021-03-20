package main

import (
	"sort"
)

/*
	有点像无限背包and有限背包的区别
	改个循环顺序
*/

func sliceEquals2(a, b []int) bool {
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

func combinationSum2(candidates []int, target int) [][]int {
	m := make(map[int][][]int, target+2)

	for _, j := range candidates {
		for i := target; i >= 1; i-- {
			if i == j {
				arr := []int{j}
				insert := true
				for _, p := range m[i] {
					if sliceEquals2(arr, p) {
						insert = false
						break
					}
				}
				if insert {
					m[i] = append(m[i], arr)
				}
			} else if i > j {
				for _, k := range m[i-j] {
					kk := make([]int, len(k))
					copy(kk, k)
					kk = append(kk, j)
					sort.Ints(kk)
					insert := true
					for _, p := range m[i] {
						if sliceEquals2(kk, p) {
							insert = false
							break
						}
					}
					if insert {
						m[i] = append(m[i], kk)
					}
				}
			}
			// fmt.Println(i, ":", m[i])
		}
	}

	return m[target]
}

// func main() {
// 	candidates := []int{2, 5, 2, 1, 2}
// 	target := 5
// 	fmt.Println(combinationSum2(candidates, target))
// }
