package main

func sumNumbers(root *TreeNode) int {
	res := sumNumbersRecur(root, 0)
	return res
}

func sumNumbersRecur(root *TreeNode, nowSum int) int {
	if root == nil {
		return 0
	}

	nowSum = nowSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return nowSum
	}

	return sumNumbersRecur(root.Left, nowSum) + sumNumbersRecur(root.Right, nowSum)
}
