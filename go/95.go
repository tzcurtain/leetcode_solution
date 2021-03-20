package main

func generateTrees(n int) []*TreeNode {
	tmp := make([]int, n)
	for i := range tmp {
		tmp[i] = i + 1
	}

	var res []*TreeNode
	for i := range tmp {
		nowRoot := new(TreeNode)
		nowRoot.Val = tmp[i]
		nowRoot.Left, nowRoot.Right = nil, nil
		// generateTreesRecursive(nowRoot, tmp[:i], tmp[i+1:], res))
	}
	return res
}
