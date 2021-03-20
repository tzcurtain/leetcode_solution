package main

func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			mostRight := root.Left
			for mostRight.Right != nil {
				mostRight = mostRight.Right
			}
			mostRight.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}
