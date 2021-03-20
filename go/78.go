package main

/*
	略微修改77题的函数即可供使用
*/

func subsets(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	resarr := make([]int, n)
	for i := 0; i < n; i++ {
		resarr[i] = i + 1
	}
	for i := 0; i < n; i++ {
		resarr[0], resarr[i] = resarr[i], resarr[0]
		subsetDfs(n, 1, &res, &resarr)
		resarr[0], resarr[i] = resarr[i], resarr[0]
	}
	res = append(res, []int{})
	return res
}

func subsetDfs(n, now int, res *[][]int, resarr *[]int) {
	tmpArr := make([]int, now)
	copy(tmpArr, (*resarr)[:now])
	*res = append(*res, tmpArr)
	if now == n {
		return
	}
	for i := now; i < n; i++ {
		if (*resarr)[i] < (*resarr)[now-1] { // 去重
			continue
		}
		(*resarr)[now], (*resarr)[i] = (*resarr)[i], (*resarr)[now]
		subsetDfs(n, now+1, res, resarr)
		(*resarr)[now], (*resarr)[i] = (*resarr)[i], (*resarr)[now]
	}
}
