package main

/*
	dfs解决
*/
func permute(nums []int) [][]int {
	res := make([][]int, 1)
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		loopPermute(nums, 1, &res)
		nums[0], nums[i] = nums[i], nums[0]
	}
	return res[1:]
}

func loopPermute(nums []int, pos int, res *[][]int) {
	if pos == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := pos; i < len(nums); i++ {
		nums[pos], nums[i] = nums[i], nums[pos]
		loopPermute(nums, pos+1, res)
		nums[pos], nums[i] = nums[i], nums[pos]
	}
}
