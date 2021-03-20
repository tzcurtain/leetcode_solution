package main

/*
	similar to q105
*/

var post int

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i := range inorder {
		m[inorder[i]] = i
	}
	post = len(inorder) - 1
	return buildTreeRecur(m, postorder, 0, len(inorder))
}

func buildTreeRecur(m map[int]int, postorder []int, l, r int) *TreeNode {
	if l == r {
		return nil
	}

	rootVal := postorder[post]
	post--
	root := new(TreeNode)
	root.Val = rootVal
	index := m[rootVal]
	root.Right = buildTreeRecur(m, postorder, index+1, r)
	root.Left = buildTreeRecur(m, postorder, l, index)
	return root
}
