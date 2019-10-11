package main

func generateTrees(n int) []*TreeNode {
	tmp := make([]int, n)
	for i := range tmp {
		tmp[i] = i + 1
	}

	var res []*TreeNode

	generateTreesRecursive()

	return res
}
