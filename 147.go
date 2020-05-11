package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	for head != nil && head.Next != nil {
		if head.Val <= head.Next.Val {
			head = head.Next
			continue
		}
		pre := dummy
		for pre.Next.Val < head.Val {
			pre = pre.Next
		}
		curr := head.Next
		head.Next = curr.Next
		curr.Next = pre.Next
		pre.Next = curr
	}

	return dummy.Next
}
