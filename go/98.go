package main

func isValidBSTRecur(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}

	if !isValidBSTRecur(root.Left, minVal, root.Val) {
		return false
	}
	if !isValidBSTRecur(root.Right, root.Val, maxVal) {
		return false
	}

	return true
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTRecur(root, IntMin, IntMax)
}
