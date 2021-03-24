#include "stdc++.h"
using namespace std;

// 678:有效的括号字符串

class Solution {
public:
    bool checkValidString(string s) {
        vector<int> star, leftP;
        int sLen = s.length();
        for (int i = 0; i < sLen; i ++) {
            if (s[i] == '(') {
                leftP.push_back(i);
            } else if (s[i] == ')') {
                if (leftP.empty()) {
                    if (star.empty()) {
                        return false;
                    }
                    star.pop_back();
                } else {
                    leftP.pop_back();
                }
            } else {
                star.push_back(i);
            }
        }

        if (!leftP.empty()) {
            int j = 0;
            for (int i = 0; i < leftP.size(); i++) {
                int now = leftP[i];
                while (j < star.size() && star[j] < now) {
                    ++j;
                }
                if (j == star.size()) {
                    return false;
                }
                j ++;
            }
        }
        return true;    
    }
};