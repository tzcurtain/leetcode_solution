#include <bits/stdc++.h> // for conveience
using namespace std;

/*
    找规律，计算位置
 */

class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows == 1) return s;
        int mod = 2*numRows-2,len = s.length();
        stringstream ss;
        for (int j=1;j<=numRows;++j){
            if (j == 1 || j == numRows){
                for (int i=0;mod*i+j<=len;++i){
                    ss << s[mod*i+j-1];
                }
            }
            else {
                for (int i=0;mod*i+j<=len;++i){
                    ss << s[mod*i+j-1];
                    int other = mod*i+2*numRows-j;
                    if (other <= len){
                        ss << s[other-1];
                    }
                }
            }
        }
        return ss.str();
    }
};