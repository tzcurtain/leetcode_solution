package main

import "sort"

type twoIntArr [][]int

/* 排序后合并 */

func (a twoIntArr) Len() int {
	return len(a)
}

func (a twoIntArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a twoIntArr) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func merge(intervals [][]int) [][]int {
	return mergeTwoIntArr(twoIntArr(intervals))
}

func mergeTwoIntArr(intervals twoIntArr) [][]int {
	var res [][]int
	now := 0
	sort.Sort(intervals)
	for i := range intervals {
		if now == 0 {
			res = append(res, intervals[i])
			now++
		} else {
			if intervals[i][0] <= res[now-1][1] {
				res[now-1][1] = max(res[now-1][1], intervals[i][1])
			} else {
				res = append(res, intervals[i])
				now++
			}
		}
	}

	return res
}
