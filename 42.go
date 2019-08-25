package main

/*
	每个位置i可以接的雨水：
		左、右边最高(包括自身)的柱子的高度a，b
		min(a,b) - height[i]
*/



func trap(height []int) int {
	lmaxi := make([]int, len(height)+2)
	rmaxi := make([]int, len(height)+2)
	nowmax, nowmaxi := -1, -1
	for i := 0; i < len(height); i++ {
		if height[i] >= nowmax {
			nowmax, nowmaxi = height[i], i
		}
		lmaxi[i] = nowmaxi
	}
	nowmax, nowmaxi = -1, -1
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] >= nowmax {
			nowmax, nowmaxi = height[i], i
		}
		rmaxi[i] = nowmaxi
	}

	res := 0
	for i := 1; i < len(height)-1; i++ {
		res += min(height[lmaxi[i]], height[rmaxi[i]]) - height[i]
	}

	return res
}
