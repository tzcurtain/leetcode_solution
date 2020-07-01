package main

func findMin(nums []int) int {
	return findMinRecur(nums, 0, len(nums)-1)
}

func findMinRecur(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	if l == r-1 {
		return min(nums[l], nums[r])
	}

	mid := (l + r) / 2
	if nums[l] < nums[r] {
		return findMinRecur(nums, l, mid)
	} else if nums[l] > nums[r] {
		if nums[mid] <= nums[r] {
			return findMinRecur(nums, l, mid)
		} else {
			return findMinRecur(nums, mid, r)
		}
	} else {
		return min(findMinRecur(nums, l, mid), findMinRecur(nums, mid+1, r))
	}
}
