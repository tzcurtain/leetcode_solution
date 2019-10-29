package main

func minDepth(root *TreeNode) int {
	return minDepthRecur(root, 1)
}

func minDepthRecur(root *TreeNode, level int) int {
	if root == nil {
		return level - 1
	}
	if root.Right == nil {
		return minDepthRecur(root.Left, level+1)
	}
	if root.Left == nil {
		return minDepthRecur(root.Right, level+1)
	}

	return min(minDepthRecur(root.Left, level+1), minDepthRecur(root.Right, level+1))
}
