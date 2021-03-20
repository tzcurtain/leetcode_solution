package main

/*
	a xor a = 0
*/
func singleNumber(nums []int) int {
	// if nums == nil || len(nums) == 0 {
	// 	return 0
	// }
	res := nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		res = res ^ nums[i]
	}
	return res
}
