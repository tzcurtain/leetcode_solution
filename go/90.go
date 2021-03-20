package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		nums[0], nums[i] = nums[i], nums[0]
		subsetDfsWithDup(n, 1, &res, &nums)
		nums[0], nums[i] = nums[i], nums[0]
	}
	res = append(res, []int{})
	return res
}

func subsetDfsWithDup(n, now int, res *[][]int, resarr *[]int) {
	tmpArr := make([]int, now)
	copy(tmpArr, (*resarr)[:now])
	*res = append(*res, tmpArr)
	if now == n {
		return
	}
	for i := now; i < n; i++ {
		if (*resarr)[i] == (*resarr)[now-1] { // 去重
			continue
		}
		(*resarr)[now], (*resarr)[i] = (*resarr)[i], (*resarr)[now]
		subsetDfsWithDup(n, now+1, res, resarr)
		(*resarr)[now], (*resarr)[i] = (*resarr)[i], (*resarr)[now]
	}
}
