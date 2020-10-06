package main

// Recursive Method

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	len := 0
	h := head
	//求出节点数
	for h != nil {
		len++
		h = h.Next
	}

	reorderListHelper(head, len)
}

func reorderListHelper(head *ListNode, len int) *ListNode {
	var outTail *ListNode
	if len == 1 {
		outTail = head.Next
		head.Next = nil
		return outTail
	}
	if len == 2 {
		outTail = head.Next.Next
		head.Next.Next = nil
		return outTail
	}
	//得到对应的尾节点，并且将头结点和尾节点之间的链表通过递归处理
	tail := reorderListHelper(head.Next, len-2)
	subHead := head.Next //中间链表的头结点
	head.Next = tail
	outTail = tail.Next //上一层 head 对应的 tail
	tail.Next = subHead
	return outTail
}
