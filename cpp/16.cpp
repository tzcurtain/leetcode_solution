#include <bits/stdc++.h>
using namespace std;

/*
    15题微改下
 */
class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        sort(nums.begin(),nums.end());
        int len = nums.size(), res = nums[0] + nums[1] + nums[2];
        for (int it = 0; it < len; ++it){
            if (it > 0 && nums[it] == nums[it-1]) continue;
            int it2 = it + 1, it3 = len-1;
            while (it2 < it3){
                int now = nums[it2] + nums[it3] + nums[it];
                if (abs(res-target) > abs(now-target)){
                    res = now;
                }
                if(now == target){
                    return res;  
                }
                else if(now > target){
                    --it3;
                }
                else
                    ++it2;
            }            
        }   
        return res; 
    }
};