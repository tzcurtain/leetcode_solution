#include <bits/stdc++.h>
using namespace std;

/* 
    分多种情况进行讨论，判断target应该落在l,mid,r中的哪个位置
    然后二分找结果
*/
class Solution {
public:
    int search(vector<int>& nums, int target, int l, int r) {
        if (nums[l] == target) return l;
        if (nums[r] == target) return r;
        if (r == l + 1 || r == l){    
            return -1;
        }
        int mid = (l + r) / 2;
        if (target == nums[mid]) return mid;
        if (target > nums[mid]){
            if (nums[l] > nums[mid] && nums[mid] < nums[r]){
                int res1 = search(nums,target,mid+1,r);
                int res2 = search(nums,target,l,mid-1);
                return (res1 != -1)? res1: res2;
            }
            else
                return search(nums,target,mid+1,r);
        }
        else {
            if (nums[mid] > nums[l] && nums[mid] > nums[r]){
                int res1 = search(nums,target,mid+1,r);
                int res2 = search(nums,target,l,mid-1);
                return (res1 != -1)? res1: res2;
            }
            else 
                return search(nums,target,l,mid-1);
        }

    }
    int search(vector<int>& nums, int target) {
        if (nums.size() == 0) return -1;
        return search(nums, target, 0, nums.size()-1);
    }
};