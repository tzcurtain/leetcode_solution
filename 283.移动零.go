/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	z, n := 0, len(nums)

	for z < n && nums[z] != 0 {
		z++
	}

	notz := z + 1

	for notz < n && nums[notz] == 0 {
		notz++
	}

	if notz == n {
		return
	}

	for i := z; i < n; i++ {
		if nums[i] == 0 {
			nums[notz], nums[i] = nums[i], nums[notz]
			notz++
		}
		for notz < n && nums[notz] == 0 {
			// fmt.Println(notz)
			notz++
		}
		if notz == n {
			return
		}
		// fmt.Println(nums, i, n, notz)
	}
}

// @lc code=end

