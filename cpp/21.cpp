#include <bits/stdc++.h>
using namespace std;
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode* dummy = new ListNode(0);
        ListNode* now = dummy;
        while (l1 != NULL && l2 != NULL){
            if (l1->val < l2->val){
                now->next = l1;
                l1 = l1->next;
            }
            else {
                now->next = l2; 
                l2 = l2->next;
            }
            now = now->next;
        }
        now->next = (l1 != NULL)? l1 : l2;
        return dummy->next;
    }
};