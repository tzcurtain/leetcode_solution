#include <bits/stdc++.h> // for conveience
using namespace std;
/*
    dfs即可
 */
class Solution {
private:
    char phone[10][5];
public:
    vector<string> letterCombinations(string digits) {
        vector<string> res;
        if (digits.length() == 0) return res;
        char now = 'a';
        for (int i=2;i<=9;++i){
            int bound = (i==7 || i == 9)? 4 : 3;
            phone[i][0] = bound;
            // cout << i << " " << bound << endl;
            for (int j=1;j<=bound;++j){
                phone[i][j] = now;
                ++now;
            }
        }
        string tmp;
        dfs(digits,0,res,tmp);
        return res;
    }
    void dfs(const string& digits,int pos,vector<string>& res,const string& tmp){
        if (pos == digits.length()){
            res.push_back(tmp);
            return;
        }
        int nowdigit = digits[pos] - '0';
        for (int i=1;i<=phone[nowdigit][0];++i){
            //cout << i << " " << nowdigit << " " << (int)phone[nowdigit][0] << endl;
            dfs(digits,pos+1,res,tmp + phone[nowdigit][i]);
        }
    }
};