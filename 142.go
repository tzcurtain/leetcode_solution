package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	t1, t2 := head, head
	for t1 != nil && t2 != nil {
		if t2.Next == nil || t1.Next == nil || t1.Next.Next == nil {
			return &ListNode{-1, nil}
		}
		t1 = t1.Next.Next
		t2 = t2.Next
		if t1 == t2 {
			t1 = head
			for t1 != t2 {
				t1 = t1.Next
				t2 = t2.Next
			}
			return t1
		}
	}
	return &ListNode{-1, nil}
}
