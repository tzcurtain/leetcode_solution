#include <bits/stdc++.h>
using namespace std;

/*
    栈的应用
 */
class Solution {
private:
    stack<char> st;
public:
    bool isValid(string s) {
        int len = s.length();
        for (int i=0;i<len;++i){
            if (s[i] == '(' || s[i] == '[' || s[i] == '{'){
                st.push(s[i]);
            }
            else {
                if (st.size() == 0 ||
                    (s[i] == ')' && st.top() != '(') || 
                    (s[i] == '}' && st.top() != '{') ||
                    (s[i] == ']' && st.top() != '['))
                    return false;
                st.pop();
            }
        }    
        return st.size() == 0; 
    }
};