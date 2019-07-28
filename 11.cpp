#include <bits/stdc++.h> // for conveience
using namespace std;
/*
    Two-Pointer

    area(i,j) (height[i] > height[j]，选择j--){
        这种情况下，我们放弃了area(i+1,j) - area(j-1,j)状态的计算
        假设有 i+1=<t<=j-1使得area(t,j) > area(i,j),
            (j-t) * min(h[j],h[t]) > (j-i) * min(h[j],h[i])
            ∵ h[i] > h[j] min(h[j],h[t])_max = h[j] (h[j] == h[t])
            ∴ (j-t)*h[j] > (j-i)*h[j]
            ∴ j-t > j-i => i > t 不成立
            ∴ 不成立
    }
 */
class Solution {
public:
    int maxArea(vector<int>& height) {
        int i = 0, j = height.size() - 1;
        int res = 0;
        while (i < j){
            res = max((j-i)*min(height[i],height[j]),res);
            if (height[i] > height[j]) j--;
            else i++;
        }
        return res;
    }
};