package main

func maxDepth(root *TreeNode) int {
	return maxDepthRecur(root, 1)
}

func maxDepthRecur(root *TreeNode, level int) int{
	if root == nil {
		return level - 1
	}
	return max(maxDepthRecur(root.Left, level+1), maxDepthRecur(root.Right, level+1))
}

