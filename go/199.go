package main

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for true {
		now := -1
		var nextqueue []*TreeNode
		pos := 0
		for pos < len(queue) {
			nowNode := queue[pos]
			if nowNode.Left != nil {
				nextqueue = append(nextqueue, nowNode.Left)
			}
			if nowNode.Right != nil {
				nextqueue = append(nextqueue, nowNode.Right)
			}
			now = nowNode.Val
			pos++
		}
		if now == -1 {
			break
		}
		res = append(res, now)
		queue = nextqueue
	}
	return res
}
