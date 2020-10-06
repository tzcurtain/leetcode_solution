package main

func majorityElement(nums []int) int {
	count := 0
	res := nums[0]
	for i := range nums {
		if count == 0 {
			res = nums[i]
		}
		if nums[i] == res {
			count++
		} else {
			count--
		}
	}

	return res
}
