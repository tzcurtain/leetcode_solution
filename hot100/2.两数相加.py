#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = ListNode()
        res = head
        x = 0 # 进位
        while l1 is not None and l2 is not None:
            head.next = ListNode()
            head = head.next      
            head.val = l1.val + l2.val + x 
            if head.val >= 10:
                head.val -= 10
                x = 1
            else:
                x = 0
            l1 = l1.next 
            l2 = l2.next 

        if l1 is not None:
            while l1 is not None:
                head.next = ListNode() 
                head = head.next
                head.val = x + l1.val 
                if head.val >= 10:
                    head.val -= 10
                    x = 1
                else:
                    x = 0
                l1 = l1.next        
        else:
            while l2 is not None:
                head.next = ListNode() 
                head = head.next
                head.val = x + l2.val 
                if head.val >= 10:
                    head.val -= 10
                    x = 1
                else:
                    x = 0
                l2 = l2.next

        if x != 0: # 处理进位
            head.next = ListNode() 
            head = head.next
            head.val = x 

        return res.next




# @lc code=end

