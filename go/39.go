package main

import (
	"sort"
)

func sliceEqual(a, b []int) bool {
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

func combinationSum(candidates []int, target int) [][]int {
	m := make(map[int][][]int, target+2)

	for i := 1; i <= target; i++ {
		for _, j := range candidates {

			if i == j {
				arr := []int{j}
				m[i] = append(m[i], arr)
			} else if i > j {
				for _, k := range m[i-j] {
					kk := make([]int, len(k))
					copy(kk, k)
					kk = append(kk, j)
					sort.Ints(kk)
					insert := true
					for _, p := range m[i] {
						if sliceEqual(kk, p) {
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
// 	candidates := []int{7, 3, 2}
// 	target := 18
// 	fmt.Println(combinationSum(candidates, target))
// }
