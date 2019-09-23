package main

func search(nums []int, target int) bool {
	return len(nums) != 0 && searchRotation(nums, target, 0, len(nums)-1)
}

func searchRotation(nums []int, target, l, r int) bool {
	if nums[l] == target || nums[r] == target {
		return true
	}
	if r == l+1 || r == l {
		return false
	}
	mid := (l + r) / 2
	if target == nums[mid] {
		return searchRotation(nums, target, l, mid)
	}
	if target > nums[mid] {
		if nums[l] >= nums[mid] && nums[mid] <= nums[r] {
			return searchRotation(nums, target, mid+1, r) || searchRotation(nums, target, l, mid-1)
		} else {
			return searchRotation(nums, target, mid+1, r)
		}
	} else {
		if nums[mid] >= nums[l] && nums[mid] >= nums[r] {
			return searchRotation(nums, target, mid+1, r) || searchRotation(nums, target, l, mid-1)
		} else {
			return searchRotation(nums, target, l, mid-1)
		}
	}

}
