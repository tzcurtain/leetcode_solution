#include <bits/stdc++.h> // for conveience
using namespace std;

/*
    中心扩展：
        找以s[i]为中心的回文子串的最长长度
 */
class Solution {
public:
    string longestPalindrome(string s) {
        int len = s.length();
        int res = 0,pos = 0;
        for (int i = 0;i < len; ++i){
            int nowres = max(extended_from_center(s,i,i),extended_from_center(s,i,i+1));
           // cout << i << " " << nowres << endl;
            if (nowres > res){
                res = nowres; pos = i - (res-1)/2;
            }
        }
        return s.substr(pos,res);
    }
    int extended_from_center(const string& s,int left,int right){
        int len = s.length();
        while (left >= 0 && right < len && s[left] == s[right]){
            --left; ++right;
        }
        return right - left - 1;
    }
};

int main(){
    Solution s;
    cout << s.longestPalindrome("cbbd");
}