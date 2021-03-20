package main

func partition(head *ListNode, x int) *ListNode {
	var l, r, resl, resr *ListNode
	for head != nil {
		if head.Val < x {
			if l != nil {
				l.Next = head
				l = head
			} else {
				l = head
				resl = head
			}
		} else {
			if r != nil {
				r.Next = head
				r = head
			} else {
				r = head
				resr = head
			}
		}
		head = head.Next
	}

	if l != nil {
		l.Next = resr
	} else {
		resl = resr
	}

	if r != nil {
		r.Next = nil
	}

	return resl
}
