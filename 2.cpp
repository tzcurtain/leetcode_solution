#include <bits/stdc++.h> // for conveience
using namespace std;
/**
 * Definition for singly-linked list.
 **/
struct ListNode {
 int val;
 ListNode *next;
 ListNode(int x) : val(x), next(NULL) {}
};


class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* dummy = new ListNode(0);
        ListNode* now = dummy;
        int x = 0;
        while (l1 != NULL && l2 != NULL){
            ListNode* aft = new ListNode((l1->val + l2->val + x) % 10);
            x = (l1->val + l2->val + x) / 10;
            l1 = l1->next; 
            l2 = l2->next;
            now->next = aft;
            now = aft; 
        }    
        while (l1 != NULL){
            ListNode* aft = new ListNode((l1->val + x) % 10);
            x = (l1->val + x) / 10;
            l1 = l1->next; 
            now->next = aft;
            now = aft;
        }
        while (l2 != NULL){
            ListNode* aft = new ListNode((l2->val + x) % 10);
            x = (l2->val + x) / 10;
            l2 = l2->next; 
            now->next = aft;
            now = aft; 
        }
        if (x != 0){
            ListNode* aft = new ListNode(x);
            now->next = aft;
        }
        return dummy->next;
    }
};