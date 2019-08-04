#include <bits/stdc++.h>
using namespace std;

/*
    固定两个数成一个数，转换成三数和做
 */
class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums,int target) {
        sort(nums.begin(),nums.end());
        int len = nums.size();
        vector<vector<int>> res;
        for (int it = 0; it < len; ++it){
            if (nums[it] > target) break;
            if (it > 0 && nums[it] == nums[it-1]) continue;
            for (int it2 = it+1; it2 < len; ++it2){
                if (nums[it] + nums[it2] > target) break;
                if (it2 > it + 1 && nums[it2] == nums[it2-1]) continue;
                int it3 = it2+1,it4 = len-1;
                while (it3 < it4){
                    int now = nums[it] + nums[it2] + nums[it3] + nums[it4];
                    if (now == target){
                        res.push_back({nums[it],nums[it2],nums[it3],nums[it4]});
                        while(it3<it4 && nums[it3+1] == nums[it3]) ++it3; //跳过重复元素
                        while(it3<it4 && nums[it4-1] == nums[it4]) --it4;  //跳过重复元素
                        ++it3, --it4;
                    }
                    else if(now > target)
                        --it4;
                    else {
                        ++it3;
                    }
                }
                
            }
        }   
        return res; 
    }
};