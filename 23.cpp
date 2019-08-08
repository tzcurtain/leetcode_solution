#include <bits/stdc++.h>
using namespace std;
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        int len = lists.size();
        if (len == 0) return NULL;
        return mergeKLists(lists, 0, len-1);
    }
    ListNode* mergeKLists(vector<ListNode*>& lists, int left, int right){

        if (left == right - 1){
            return mergeLists(lists[left],lists[right]);
        }
        else if (left == right) {
            return lists[left];
        }         
        else {
            int mid = (left + right) / 2;
            ListNode* l = mergeKLists(lists,left,mid);
            cout << l->val;
            ListNode* r = mergeKLists(lists,mid+1,right);
            cout << r->val;
            return mergeLists(l,r);
        }
    }
    ListNode* mergeLists(ListNode* l,ListNode* r){
        ListNode* dummy = new ListNode(0);
        ListNode* now = dummy;
        while (l != NULL && r != NULL){
            if (l->val < r->val){
                now->next = l;
                l = l->next;
            }
            else {
                now->next = r;
                r = r->next;
            }
            now = now->next;
        }
        if (l != NULL){
            now->next = l;
        }
        else{
            now->next = r;
        }
        return dummy->next;
    }
};

int main(){
    return 0;
}