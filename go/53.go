package main

/*
	分治：最大子序列和的要不在左半边，要不在右半边，要不就是穿过中间的序列
*/

func maxSubArray(nums []int) int {
	return maxSubArrayDiv(nums, 0, len(nums)-1)
}

func maxSubArrayDiv(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	mid := (l + r) / 2

	maxL, maxR := nums[mid], nums[mid]
	now := maxL
	for i := mid - 1; i >= l; i-- {
		now += nums[i]
		maxL = max(maxL, now)
	}
	now = maxR
	for i := mid + 1; i <= r; i++ {
		now += nums[i]
		maxR = max(maxR, now)
	}

	return max(maxSubArrayDiv(nums, l, mid), max(maxSubArrayDiv(nums, mid+1, r), maxL+maxR-nums[mid]))
}
