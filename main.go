package main

import "fmt"

func main() {
	l1, l2, l3 := new(ListNode), new(ListNode), new(ListNode)
	l1.Val, l2.Val, l3.Val = 0, 1, 2
	l1.Next, l2.Next = l2, l3
	fmt.Print(rotateRight(l1, 4))

}
