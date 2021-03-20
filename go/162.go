package main

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		if l == r-1 {
			return l
		}
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
