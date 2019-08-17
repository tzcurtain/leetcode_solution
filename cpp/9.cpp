#include <bits/stdc++.h> // for conveience
using namespace std;

/*
 */
class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0) return false;
        string s = to_string(x);
        int i = 0,len = s.length();
        bool res = true;
        while (i < len - i - 1){
            if (s[i] != s[len-1-i]){
                res = false;
                break;
            }
            ++i;
        }   
        return res;
    }
};