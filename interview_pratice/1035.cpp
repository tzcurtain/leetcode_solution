#include "stdc++.h"
using namespace std;

// 1035: 不相交的线

// 本质LCS
class Solution {
public:
    int maxUncrossedLines(vector<int>& A, vector<int>& B) {
        int aLen = A.size(), bLen = B.size();
        int** dp = new int*[aLen + 1];
        for (int i = 0; i < aLen + 1; ++i) {
            dp[i] = new int[bLen + 1];
            memset(dp[i], 0, sizeof(int)*(bLen+1));
        }

        int ans = 0;
        for (int i = 1; i <= aLen; ++i) {
            for (int j = 1; j <= bLen; ++j) {
                if (A[i-1] == B[j-1]) {
                    dp[i][j] = dp[i-1][j-1] + 1;
                } else {
                    dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
                }
                ans = max(dp[i][j], ans);
            }
        }
        return ans;
    }
};