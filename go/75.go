package main

/*
	p0, p2, cur 三指针解决
*/

func sortColors(nums []int) {
	p0, p2, cur := 0, len(nums)-1, 0
	for cur <= p2 {
		switch nums[cur] {
		case 0:
			nums[cur], nums[p0] = nums[p0], nums[cur]
			p0++
			cur++
		case 1:
			cur++
		case 2:
			nums[cur], nums[p2] = nums[p2], nums[cur]
			p2--
		}

	}
}
