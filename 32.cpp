#include <bits/stdc++.h>
using namespace std;


class Solution {
public:
    int longestValidParentheses(string s) {
        int len = s.length();
        int l = 0,r = 0,res = 0;
        for (int i=0;i<len;++i){
            if (s[i] == '(') ++l;
            else if (s[i] == ')') ++r;
            if (l == r) res = max(res,l);
            if (r > l){
                l = 0,r = 0;
            }
        }
        l = 0,r = 0;
        for (int i=len-1;i>=0;--i){
            if (s[i] == '(') ++l;
            else if (s[i] == ')') ++r;
            if (l == r) res = max(res,l);
            if (l > r){
                l = 0,r = 0;
            }
        }
        return 2*res;
    }
};