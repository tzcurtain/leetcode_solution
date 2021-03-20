package main

import "fmt"

/*
	一个一个处理 注意边界
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	if len(intervals) == 0 || intervals[0][0] > newInterval[1] {
		res = append(res, newInterval)
		if len(intervals) == 0 {
			return res
		}
		if intervals[0][0] > newInterval[1] {
			res = append(res, intervals...)
		}
		return res
	}
	if newInterval[0] > intervals[len(intervals)-1][1] {
		res = append(res, intervals...)
		res = append(res, newInterval)
		return res
	}

	insert := false
	for i := 0; i < len(intervals); i++ {
		fmt.Println(i)
		if !insert && intervals[i][0] > newInterval[1] {
			res = append(res, newInterval)
			insert = true
		}
		res = append(res, intervals[i])
		if insert || notInvolve(intervals[i], newInterval) {
			continue
		} else {
			insert = true
			res[len(res)-1][0] = min(newInterval[0], res[len(res)-1][0])
			res[len(res)-1][1] = max(newInterval[1], res[len(res)-1][1])
			for i+1 < len(intervals) && !notInvolve(res[len(res)-1], intervals[i+1]) {
				i++
				res[len(res)-1][0] = min(res[len(res)-1][0], intervals[i][0])
				res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
			}
		}
	}
	return res
}

func notInvolve(a, b []int) bool {
	// fmt.Println(a, b)
	return a[0] > b[1] || a[1] < b[0]
}
