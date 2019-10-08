package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	befhead := dummy
	for i := 1; i < m; i++ {
		if i == m-1 {
			befhead = head
		}
		head = head.Next
	}
	nowhead := head
	nowtail := head
	var bef *ListNode
	for i := 1; i <= n-m; i++ {
		// fmt.Println(head.Val)
		if i == n-m {
			nowtail = head
		}
		tmp := head.Next
		head.Next = bef
		bef = head
		head = tmp
	}
	aftertail := head

	befhead.Next = nowtail
	nowhead.Next = aftertail

	return dummy.Next
}
