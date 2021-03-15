/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// jump
	oddtail := head
	eventail := head.Next
	evenhead := eventail
	if eventail == nil {
		return head
	}
	now := eventail.Next
	if now == nil {
		return head
	}
	odd := true
	for now != nil {
		if odd {
			oddtail.Next = now
			oddtail = now
		} else {
			eventail.Next = now
			eventail = now
		}
		odd = !odd
		nextNow := now.Next
		now.Next = nil
		now = nextNow
	}
	oddtail.Next = evenhead
	eventail.Next = nil
	return head
}

// @lc code=end

