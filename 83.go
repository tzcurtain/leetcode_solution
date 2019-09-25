package main

/*
	easier version of 82
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	res := head
	now := head.Next
	for now != nil {
		if now.Val == res.Val {
			now = now.Next
		} else {
			res.Next = now
			res = now
			now = now.Next
		}
	}
	res.Next = nil
	return head
}
