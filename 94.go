package main

/*
	递归
*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

/*
	迭代
*/
func inorderTraversalNonRecursive(root *TreeNode) []int {
	var res []int
	st := make([]*TreeNode, 1000)
	stLen := 0
	tmp := root
	for tmp != nil || stLen > 0 {
		for tmp != nil {
			stLen++
			st[stLen] = tmp
			tmp = tmp.Left
		}
		if stLen > 0 {
			tmp = st[stLen]
			res = append(res, tmp.Val)
			stLen--
			tmp = tmp.Right
		}
	}

	return res
}
