package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var res, tmp *ListNode
	now := head
	len := 1
	for ; now.Next != nil; len++ {
		now = now.Next
	}
	now = head
	nowlen := 0
	k = k % len
	if k == 0 {
		return head
	}
	for ; now != nil; nowlen++ {
		if nowlen == len-k {
			res = now
		}
		tmp = now
		now = now.Next
		if nowlen == len-k-1 {
			tmp.Next = nil
		}
	}
	tmp.Next = head
	return res
}
