package main

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var f [2][4]int // 0 负绝对值最大 1 正值最大 2 不包含i的负绝对值最大 3 不包含i的正最大
	res := 0
	if nums[0] > 0 {
		f[0][1] = nums[0]
	} else {
		f[0][0] = nums[0]
	}

	now := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			if abs(f[1-now][1]*nums[i]) > abs(nums[i]) {
				f[now][0] = f[1-now][1] * nums[i]
			} else {
				f[now][0] = nums[i]
			}
			f[now][1] = f[1-now][0] * nums[i]
		} else {
			if f[1-now][1] != 0 {
				f[now][1] = f[1-now][1] * nums[i]
			} else {
				f[now][1] = nums[i]
			}
			f[now][0] = f[1-now][0] * nums[i]
		}

		if abs(f[1-now][0]) > abs(f[1-now][2]) {
			f[now][2] = f[1-now][0]
		} else {
			f[now][2] = f[1-now][2]
		}

		f[now][3] = max(f[1-now][1], f[1-now][3])

		res = max(res, max(f[now][0], max(f[now][1], max(f[now][2], f[now][3]))))
		now = 1 - now
	}

	return res
}
