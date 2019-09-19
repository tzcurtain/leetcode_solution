package main

func combine(n int, k int) [][]int {
	var res [][]int
	resarr := make([]int, n)
	for i := 0; i < n; i++ {
		resarr[i] = i + 1
	}
	for i := 0; i < n; i++ {
		resarr[0], resarr[i] = resarr[i], resarr[0]
		combineDfs(n, k, 1, &res, &resarr)
		resarr[0], resarr[i] = resarr[i], resarr[0]
	}
	return res
}

func combineDfs(n, k, now int, res *[][]int, resarr *[]int) {
	if now == k {
		tmpArr := make([]int, k)
		copy(tmpArr, (*resarr)[:k])
		*res = append(*res, tmpArr)
		return
	}
	for i := now; i < n; i++ {
		if (*resarr)[i] < (*resarr)[now-1] { // 去重
			continue
		}
		(*resarr)[now], (*resarr)[i] = (*resarr)[i], (*resarr)[now]
		combineDfs(n, k, now+1, res, resarr)
		(*resarr)[now], (*resarr)[i] = (*resarr)[i], (*resarr)[now]
	}
}
