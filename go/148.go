package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slowP, fastP := head, head.Next
	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}
	r := mergeSort(slowP.Next)
	slowP.Next = nil
	l := mergeSort(head)
	return mergeList(l, r)
}

func mergeList(l, r *ListNode) *ListNode {
	tmpHead := new(ListNode)
	p := tmpHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l == nil {
		p.Next = r
	} else {
		p.Next = l
	}
	return tmpHead.Next
}
