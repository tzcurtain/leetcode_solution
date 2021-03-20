package main

func reverseList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	var bef *ListNode
	for head != nil {
		dummy.Next = head
		tmp := head.Next
		head.Next = bef
		bef = head
		head = tmp
	}

	return dummy.Next
}
