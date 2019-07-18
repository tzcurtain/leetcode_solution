#include <bits/stdc++.h> // for conveience
using namespace std;

class Solution {
private:
    map<int,int> m;
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> res;
        int len = nums.size();
        for (int i=0;i < len; ++i){
            int diff = target-nums[i];
            if (m.find(diff) == m.end()){
                if (m.find(nums[i]) == m.end())
                    m[nums[i]] = i;
            }
            else {
                res.push_back(m[diff]);
                res.push_back(i);
                return res;
            }
        }
        return res;
    }
};