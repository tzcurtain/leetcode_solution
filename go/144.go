package main

func preorderTraversalRecur(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	tmp := make([]int, 1)
	tmp[0] = root.Val
	tmp = append(tmp, preorderTraversal(root.Left)...)
	tmp = append(tmp, preorderTraversal(root.Right)...)
	return tmp
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var queue []*TreeNode
	stack := make([]*TreeNode, 100)
	var res []int
	head := 0
	tail := 1
	top := -1
	queue = append(queue, root)
	for head != tail || top != -1 {

		if head == tail { // queue is empty
			queue = append(queue, stack[top])
			top--
			tail++
			continue
		}

		nowQ := queue[head]
		res = append(res, nowQ.Val)
		if nowQ.Left != nil {
			queue = append(queue, nowQ.Left)
			tail++
		}
		if nowQ.Right != nil {
			top++
			stack[top] = nowQ.Right
		}
		head++
	}

	return res
}
