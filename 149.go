package main

import "sort"

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] == points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	res := 0
	n := len(points)
	for i := 0; i < n; i++ {
		same := 0
		for j := i + 1; j < n; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				same++
				continue
			}
			total := 2
			for k := j + 1; k < n; k++ {
				if (points[j][1]-points[i][1])*(points[k][0]-points[i][0]) == (points[k][1]-points[i][1])*(points[j][0]-points[i][0]) {
					total++
				}
			}
			res = max(res, total+same)
		}
		res = max(same+1, res)
	}

	return res
}
