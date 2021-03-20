package main

// var pre int

// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	m := make(map[int]int, len(preorder))
// 	for i := range inorder {
// 		m[inorder[i]] = i
// 	}
// 	pre = 0
// 	return buildTreeRecur(m, preorder, 0, len(inorder))
// }

// func buildTreeRecur(m map[int]int, preorder []int, l, r int) *TreeNode {
// 	if l == r {
// 		return nil
// 	}

// 	rootVal := preorder[pre]
// 	pre++
// 	root := new(TreeNode)
// 	root.Val = rootVal
// 	index := m[rootVal]
// 	root.Left = buildTreeRecur(m, preorder, l, index)
// 	root.Right = buildTreeRecur(m, preorder, index+1, r)
// 	return root
// }
