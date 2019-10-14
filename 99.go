package main

/*
	记录最初起冲突和最后起冲突的位置
*/

func recoverTree(root *TreeNode) {
	stSize := 0
	var st []*TreeNode
	inorder := IntMin
	var first, last, inorderPos *TreeNode
	for stSize != 0 || root != nil {
		for root != nil {
			st = append(st, root)
			stSize++
			root = root.Left
		}
		stSize--
		root = st[stSize]
		st = st[:stSize]
		if root.Val <= inorder {
			// do sth
			if first == nil {
				first = inorderPos
			}
			last = root
		} else {
			inorder = root.Val
			inorderPos = root
		}
		root = root.Right
	}
	first.Val, last.Val = last.Val, first.Val
}
