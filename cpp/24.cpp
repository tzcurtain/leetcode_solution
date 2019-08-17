#include <bits/stdc++.h>
using namespace std;
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
/*
    use dummy to make it simple
 */
class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        ListNode* dummy = new ListNode(0);
        dummy->next = head;
        ListNode* now = dummy;
        while (now->next && now->next->next){
            ListNode* nown = now->next;
            ListNode* nownn = now->next->next;
            now->next = nownn;
            nown->next = nownn->next;
            nownn->next = nown;
            now = nown;
        }
        return dummy->next;
    }
};