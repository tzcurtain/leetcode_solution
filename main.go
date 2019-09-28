package main

import "fmt"

func main() {

	ln := new(ListNode)
	tmp := ln
	ln.Val = 1
	ln.Next = new(ListNode)
	ln = ln.Next
	ln.Val = 4
	ln.Next = new(ListNode)
	ln = ln.Next
	ln.Val = 3
	ln.Next = new(ListNode)
	ln = ln.Next
	ln.Val = 2
	ln.Next = new(ListNode)
	ln = ln.Next
	ln.Val = 5
	ln.Next = new(ListNode)
	ln = ln.Next
	ln.Val = 2

	fmt.Println(partition(tmp, 3))
}
