package main

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	// right := len(nums)-1
	if right == -1 {
		return []int{-1, -1}
	}
	for left < right {
		if nums[left] == target {
			break
		}
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	lres := left
	if nums[lres] != target {
		return []int{-1, -1}
	}
	left, right = 0, len(nums)-1
	for left < right {
		if nums[right] == target {
			break
		}
		mid := (left + right) / 2
		if nums[mid] <= target {
			if left == mid {
				right = left
				break
			}
			left = mid
		} else {
			right = mid - 1
		}
	}
	rres := right
	res := []int{lres, rres}
	return res
}

// func main(){

// }
