#include <bits/stdc++.h>
using namespace std;
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
/*
    类似24只不过变k个了
 */
class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        ListNode* dummy = new ListNode(0);
        dummy->next = head;
        ListNode* now = dummy;
        vector<ListNode*> kPtr;
        int size = 0;
        while (now){
            ListNode* bef = now;
            kPtr.clear();
            size = 0;
            while (now && size < k){
                now = now->next;
                if (!now) break;
                ++size;
                kPtr.push_back(now);
            }
            if (size == k){
                kPtr[0]->next = kPtr[k-1]->next;
                for (int i=k-1;i>=1;--i){
                    kPtr[i]->next = kPtr[i-1];
                }
                bef->next = kPtr[k-1];
                now = kPtr[0];
            }
            else break;
        }
        return dummy->next;
    }
};