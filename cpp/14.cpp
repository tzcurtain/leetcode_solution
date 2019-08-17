#include <bits/stdc++.h> // for conveience
using namespace std;
class Solution {
public:
    string longestCommonPrefix(const string& s1, const string& s2) {
        stringstream ss;
        int len = min(s1.length(),s2.length());
        int i;
        for (i=0;i<len;++i){
            if (s1[i] == s2[i]){
                ss << s1[i];        
            }
            else break;
        }
        if (i == 0) return "";
        string res;
        ss >> res;
        return res;
    }
    
    string longestCommonPrefix(vector<string>& strs) {
        int len = strs.size();
        if (len == 1) return strs[0];
        if (len == 0) return "";
        string res = longestCommonPrefix(strs[0],strs[1]);
        for (int i = 2; i < len; ++i){
            res = longestCommonPrefix(res,strs[i]);
        }
        return res;
    }
};