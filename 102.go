package main

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	levelOrderRecur(root, 0, &res)
	return res
}

func levelOrderRecur(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if *res == nil || len(*res) <= level {
		var tmp []int
		*res = append(*res, tmp)
	}

	(*res)[level] = append((*res)[level], root.Val)

	levelOrderRecur(root.Left, level+1, res)
	levelOrderRecur(root.Right, level+1, res)
}
