#include "stdc++.h"
using namespace std;

// 1249: 移除不相关的括号

class Solution {
public:
    string minRemoveToMakeValid(string s) {
        int sLen = s.length();
        int num = 0;
        stringstream ss;
        bool* ok = new bool[sLen + 1];
        memset(ok, false, sizeof(bool)*(sLen+1));
        vector<int> st;
        
        for (int i = 0; i < sLen; ++i) {
            if (s[i] == '(') {
                ok[i] = true;
                st.push_back(i);
            } else if (s[i] == ')') {
                if (st.size() == 0) { // abandon this )
                    ok[i] = true;
                    continue;
                }
                ok[st.back()] = false;
                st.pop_back();
            } 
        }

        for (int i = 0; i < sLen; ++i) {
            if (!ok[i]) {
                ss << s[i];
            }
        }

        return ss.str();
    }
};