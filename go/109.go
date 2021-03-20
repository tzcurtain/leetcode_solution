package main

/*
	空间换时间
*/

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(nums)
}
