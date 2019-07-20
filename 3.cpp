#include <bits/stdc++.h> // for conveience
using namespace std;

/*
    lastAppear[i]记载字符i上次出现在字符串的位置(从1开始计数)
    当遇到s[i]时，如果字符s[i]上次未出现，或出现在当前考虑的now子串的前面，则++now
    否则now子串从字符s[i]上次出现的位置后面一个开始重新计算
 */
class Solution {   
public:
    int lengthOfLongestSubstring(string s) {
        int lastAppear[300] = {0};  
        int len = s.length();
        int res = 0, now = 0;
        for (int i=0; i<len; ++i){
            int chIndex = s[i];
            if (lastAppear[chIndex] == 0 | lastAppear[chIndex] - 1 < i - now){
                ++now;
            }
            else {
                res = max(now,res);
                now = i - lastAppear[chIndex] + 1; 
            }
            lastAppear[chIndex] = i+1;
        }
        res = max(now,res);
        return res;
    }
};
