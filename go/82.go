package main

/*
	use dummy to get easy
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := new(ListNode)
	dummy := res
	res.Next = nil
	now := head
	for {
		if now == nil {
			res.Next = nil
			break
		}
		if now.Next == nil || now.Val != now.Next.Val {
			res.Next = now
			res = now
			now = now.Next
		} else {
			tmp := now.Next
			for tmp != nil && tmp.Val == now.Val {
				tmp = tmp.Next
			}
			now = tmp
		}
	}
	return dummy.Next
}
