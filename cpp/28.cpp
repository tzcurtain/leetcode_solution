#include <bits/stdc++.h>
using namespace std;

/*
    普通的暴力
 */
class Solution {
public:
    int strStr(string haystack, string needle) {
        int lenh = haystack.length(), lenn = needle.length();
        if (lenn == 0) return 0;
        for (int i=0;i<=lenh-lenn;++i){
            if (haystack[i] == needle[0]){
                bool match = true;
                int nowi = i+1, nown = 1;
                while (nowi < lenh && nown < lenn){
                    if (haystack[nowi] != needle[nown]){
                        match = false; break;
                    }
                    ++nowi; ++nown;
                }
                if (match && nown == lenn) return i;
            }
        }
        return -1;
    }
};