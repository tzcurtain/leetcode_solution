#include "stdc++.h"
using namespace std;

// 53:最大子序列

// O(n)算法
// class Solution {
// public:
//     int maxSubArray(vector<int>& nums) {
//         bool flag = false;
//         int res = 0, now = 0;
//         for (int i=0;i < nums.size(); i ++) {
//             now = now + nums[i];
//             if (!flag || now > res) {
//                 flag = true;
//                 res = now;
//             }
//             if (now < 0) {
//                 now = 0;
//             }
//         }
//         return res;
//     }
// };

// 分治法 建立线段树
class Solution {
public:
    struct Status{
        int lSum, rSum, mSum, iSum;
        /* lsum 表示 [l,r]内以 l 为左端点的最大子段和
           rSum 表示 [l,r]内以 r 为右端点的最大子段和
           mSum 表示 [l,r]内的最大子段和
           iSum 表示 [l,r]的区间和
        */
    };

    Status pushUp(Status l, Status r)  {
        int iSum = l.iSum + r.iSum;
        int lSum = max(l.lSum, l.iSum + r.lSum);
        int rSum = max(r.rSum, l.rSum + r.iSum);
        int mSum = max(max(l.mSum, r.mSum), l.rSum + r.lSum);
        return (Status) {lSum, rSum, mSum, iSum};
    }

    Status get(vector<int>& nums, int l, int r) {
        if (l == r) {
            return (Status) {nums[l], nums[l], nums[l], nums[l]};
        }
        int mid = (l + r) >> 1;
        Status lsub = get(nums, l, mid);
        Status rsub = get(nums, mid+1, r);
        return pushUp(lsub, rsub);
    }

    int maxSubArray(vector<int>& nums) {
        return get(nums, 0, nums.size() - 1).mSum;
    }

};