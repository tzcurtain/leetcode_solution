#include <bits/stdc++.h>
using namespace std;
/*
    dfs解决
 */
class Solution {
private:
    char tmp[1000];
    vector<string> res;
public:
    vector<string> generateParenthesis(int n) {
        dfs(1,n,0,0);
        return res;
    }

    void dfs(int i, int n,int now,int rightP){
        tmp[now] = '('; ++now;
        if (i == n){
            for (int j=1;j<=rightP+1;++j){
                tmp[now] = ')'; ++now;  
            }
            tmp[now] = '\0';
            res.push_back(string(tmp));
            return;
        }
        for (int j=rightP+1;j>=0;j--){
            int bound = now + rightP - j;
            for (int k=now;k<=bound;++k){
                tmp[k] = ')';
            }
            ++bound;
            dfs(i+1,n,bound,j);
        }
    }
};