package main

/*
	单调栈
*/

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stacks := make([]int, len(heights))
	nowStackTop, top := -1, 0
	res := 0
	heights = append(heights, -1) // last element to clear stack at last
	for i := range heights {
		if nowStackTop == -1 || heights[stacks[nowStackTop]] <= heights[i] {
			nowStackTop++
			stacks[nowStackTop] = i
		} else {
			for nowStackTop >= 0 && heights[stacks[nowStackTop]] > heights[i] {
				top = stacks[nowStackTop]
				res = max(res, (i-top)*heights[top])
				nowStackTop--
			}
			nowStackTop++
			stacks[nowStackTop] = top
			heights[top] = heights[i] // 当前入栈元素所能向左拓展的最大距离
		}
	}

	return res
}
