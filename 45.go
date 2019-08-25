package main

func jump(nums []int) int {
	lenn := len(nums)
	f := make([]int, lenn)
	for i := range f {
		f[i] = lenn
	}
	f[0] = 0
	now := nums[0]
	nowEnd := min(lenn-1, now)
	for i := 1; i <= nowEnd; i++ {
		f[i] = 1
	}
	for i := 1; i < lenn; i++ {
		if i+nums[i] <= now { // 如果小于now花的步数肯定比现在更多
			continue
		}
		nowEnd = min(i+nums[i], lenn-1)
		for j := now + 1; j <= nowEnd; j++ {
			f[j] = f[i] + 1
		}
		now = i + nums[i]
	}

	return f[lenn-1]
}
