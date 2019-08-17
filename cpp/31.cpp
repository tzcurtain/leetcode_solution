#include <bits/stdc++.h>
using namespace std;

/*
    从后面开始往前找第一个不是升序的nums[i]，找到右边比他小的nums[j]交换，
    然后翻转i+1后面的
 */
class Solution {
public:
    void nextPermutation(vector<int>& nums) {
        int i = nums.size()-2;
        while (i>=0 && nums[i+1]<=nums[i]){
            --i;
        }
        if (i >= 0){
            int j = nums.size()-1;
            while (j >= 0 && nums[j] <= nums[i]) --j;
            swap(nums[i],nums[j]);
        }
        auto it = nums.begin();
        advance(it,i+1);
        reverse(it,nums.end());
    }
};

int main(){
    vector<int> nums = {4,2,0,2,3,2,0};
    Solution s = Solution();
    s.nextPermutation(nums);
    for (int i=0;i<nums.size();++i){
        cout << nums[i] << " ";
    }
}