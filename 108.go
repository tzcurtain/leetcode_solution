package main

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	lens := len(nums)
	root := new(TreeNode)
	root.Val = nums[lens/2]
	root.Left = sortedArrayToBST(nums[:lens/2])
	root.Right = sortedArrayToBST(nums[lens/2+1:])
	return root
}

// 88ms example
// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	return buildTree(nums, 0, len(nums)-1)
// }

// func buildTree(nums []int, l, r int) *TreeNode {
// 	if l > r {
// 		return nil
// 	}
// 	m := l + (r-l)/2
// 	root := &TreeNode{Val: nums[m]}
// 	root.Left = buildTree(nums, l, m-1)
// 	root.Right = buildTree(nums, m+1, r)
// 	return root;
// }
