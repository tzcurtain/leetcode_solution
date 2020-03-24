package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	t1, t2 := head, head
	for t1 != nil && t2 != nil {
		if t2.Next == nil || t1.Next == nil || t1.Next.Next == nil {
			return false
		}
		t1 = t1.Next.Next
		t2 = t2.Next
		if t1 == t2 {
			return true
		}
	}

	return false
}
