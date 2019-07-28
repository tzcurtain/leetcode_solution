#include <bits/stdc++.h> // for conveience
using namespace std;

class Solution {
private:
    map<pair<int,int>, int> m;
public:
    bool isMatch(string s, string p) {
        return dp(0,0,s,p);
    }
    bool dp(int i,int j,const string& s,const string& p) {
        pair<int,int> pij = make_pair(i,j);
        if (m.find(pij) != m.end()) // use Memoization
            return m[pij];
        int lenp = p.length(),lens = s.length();
        if (j == lenp) return (i == lens);
        bool first = i < lens && (p[j] == s[i] || p[j] == '.');
        bool ans;
        if (j+1 < lenp && p[j+1] == '*'){
            /* 直接跳到后面 或是 使用*号匹配s一个字符 */
            ans = dp(i,j+2,s,p) || first && dp(i+1,j,s,p);
        }
        else {
            ans = first && dp(i+1,j+1,s,p);
        }
        m[pij] = ans;
        return ans;
    }
};