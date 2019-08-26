package main

import "bytes"

/*
	dfs解决
	比46多个map去重
*/
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 1)
	resm := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		loopPermuteUnique(nums, 1, &res, resm)
		nums[0], nums[i] = nums[i], nums[0]
	}
	return res[1:]
}

func makeString(nums []int) string {
	var byte bytes.Buffer
	for i := range nums {
		byte.WriteString(string(nums[i]))
	}
	return byte.String()
}

func loopPermuteUnique(nums []int, pos int, res *[][]int, resm map[string]bool) {
	if pos == len(nums) {
		numString := makeString(nums)
		_, ok := resm[numString]
		if ok == true {
			return
		}
		resm[numString] = true
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := pos; i < len(nums); i++ {
		nums[pos], nums[i] = nums[i], nums[pos]
		loopPermuteUnique(nums, pos+1, res, resm)
		nums[pos], nums[i] = nums[i], nums[pos]
	}
}
