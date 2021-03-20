package main

func isBalanced(root *TreeNode) bool {
	ok, _ := isBalancedRecur(root, 0)
	return ok
}

func isBalancedRecur(root *TreeNode, height int) (bool, int) {
	if root == nil {
		return true, height - 1
	}

	ok1, h1 := isBalancedRecur(root.Left, height+1)
	ok2, h2 := isBalancedRecur(root.Right, height+1)
	if h1 > h2+1 || h2 > h1+1 {
		return false, max(h1, h2)
	}

	return ok1 && ok2, max(h1, h2)

}
