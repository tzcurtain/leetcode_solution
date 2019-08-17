#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        sort(nums.begin(),nums.end());
        int len = nums.size();
        vector<vector<int>> res;
        for (int it = 0; it < len; ++it){
            if (nums[it] > 0) break;
            if (it > 0 && nums[it] == nums[it-1]) continue;
            int it2 = it + 1, it3 = len-1, target = -nums[it];
            while (it2 < it3){
                if(nums[it2] + nums[it3] == target){
                    res.push_back({nums[it],nums[it2],nums[it3]});
                    while(it2<it3 && nums[it2+1] == nums[it2]) ++it2; //跳过重复元素
                    while(it2<it3 && nums[it3-1] == nums[it3]) --it3;  //跳过重复元素
                    ++it2, --it3;
                }
                else if(nums[it2] + nums[it3] > target)
                    --it3;
                else
                    ++it2;
            }            
        }   
        return res; 
    }
};