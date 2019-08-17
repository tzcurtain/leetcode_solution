#include <bits/stdc++.h>
using namespace std;


struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
/*
    队列法实现
 */
// class Solution {
// private:
//     queue<ListNode*> q;
// public:
//     ListNode* removeNthFromEnd(ListNode* head, int n) {
//         ListNode* now = head;
//         while (q.size() < n){
//             q.push(now);
//             now = now->next;
//             if (now == NULL) return head->next; // n=链表长度
//         }
//         while (now->next != NULL){
//             q.pop();
//             q.push(now);
//             now = now->next;
//         }
//         ListNode* resbef = q.front();
//         if (resbef->next != NULL)
//             resbef->next = resbef->next->next;
//         return head;
//     }
// };
/*
    双指针 + dummy
 */
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* p1 = head,*p2 = head;
        for (int i=0;i<n;++i) p2 = p2->next;
        if (!p2) return head->next;
        while (p2->next){
            p1 = p1->next;
            p2 = p2->next;
        }
        p1->next = p1->next->next;
        return head;
    }
};