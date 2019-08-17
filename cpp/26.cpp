#include <bits/stdc++.h>
using namespace std;
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int size = nums.size();
        if (size == 0) return 0;
        int now = nums[0], nowi = 0;
        for (int i=1;i<size;++i){
            if (nums[i] == now) continue;
            else {
                now = nums[i];
                ++nowi;
                nums[nowi] = now; 
            }
        }
        return nowi + 1;
    }
};