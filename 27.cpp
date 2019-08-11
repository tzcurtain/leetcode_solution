#include <bits/stdc++.h>
using namespace std;
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int size = nums.size();
        int nowsize = 0;
        for (int i=0;i<size;++i){
            if (nums[i] != val){
                ++nowsize;
                nums[nowsize-1] = nums[i];
            }
        }
        return nowsize;
    }
};

