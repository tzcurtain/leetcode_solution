#include "stdc++.h"
using namespace std;

// 918:环形子数组的最大和

// A[0,i] 用Kadane算法
// A[0,i][j,n-1]的最大值转发为A整个数组减去A[i+1, j-1]

class Solution {
public:
    int maxSubarraySumCircular(vector<int>& A) {
        int n = A.size();
        if (n == 1) {
            return A[0];
        }

        vector<int> lSum;
        int ans1 = INT32_MIN, now = INT32_MIN, total = 0;
        for (int i = 0; i < n; i++){
            if (now < 0) {
                now = 0;
            }
            now += A[i];
            total += A[i];
            ans1 = max(now, ans1);
        }

        int ans2 = INT32_MAX;
        now = INT32_MAX;

        for (int i = 0; i < n-1; i++){
            if (now > 0) {
                now = 0;
            }
            now += A[i];
            ans2 = min(now, ans2);
        }
        if (ans2 != total) { // 子数组不能为空
            ans2 = total - ans2;
        }

        return max(ans1, ans2);
    }
};