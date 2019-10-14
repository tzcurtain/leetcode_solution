package main

import (
	"fmt"
	"math"
)

/*
	const
*/
const (
	Modint = 1000000007
	IntMax = int(^uint(0) >> 1)
	IntMin = ^IntMax
)

/*
ListNode type definition
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
TreeNode type definition
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* util to make linklist
@return *ListNode
*/
func makeLinkList(head *ListNode, arr []int) *ListNode {
	if len(arr) == 0 {
		return head
	}
	head.Val = arr[0]
	tmp := head
	for i := 1; i < len(arr); i++ {
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		tmp.Val = arr[i]
	}

	return head
}

func printLinkList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print("->")
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/* 迷宫方向常量 */
const (
	SpiralRight = iota
	SpiralDown
	SpiralLeft
	SpiralUp
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPrime(a int) bool {
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func isAlpha(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}
