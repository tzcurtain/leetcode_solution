package main

func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	now := dummy
	for now != nil {
		for now.Next != nil && now.Next.Val == val {
			now.Next = now.Next.Next
		}
		now = now.Next
	}

	return dummy.Next
}
