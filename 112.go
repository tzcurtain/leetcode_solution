package main

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumRecur(root, sum, 0)
}

func hasPathSumRecur(root *TreeNode, sum, nowSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil { // is leaf
		return sum == nowSum+root.Val
	}

	return hasPathSumRecur(root.Left, sum, nowSum+root.Val) ||
		hasPathSumRecur(root.Right, sum, nowSum+root.Val)
}
