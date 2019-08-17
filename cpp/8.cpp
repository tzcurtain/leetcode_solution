#include <bits/stdc++.h> // for conveience
using namespace std;

/*
    纯细节题
 */
class Solution {
public:
    int myAtoi(string str) {
        int now = 0;
        int len = str.length();
        while (now < len && str[now] == ' '){
            ++now;
        }
        if (now == len) return 0;
        bool fu = false;
        if (str[now] == '-' || str[now] == '+') { 
            fu = (str[now] == '-')? true: false; 
            ++now; 
        } 
        if (now == len || !isdigit(str[now])){ return 0;}
        long long res = 0;
        while (now < len && isdigit(str[now])){
            res = res * 10 + (str[now] - '0');
            if (res > INT_MAX && !fu) return INT_MAX;
            if (-res < INT_MIN && fu) return INT_MIN;
            ++ now;
        }
        if (fu) return -res;
        return res;
    }
};

int main(){
    Solution s;
    cout << s.myAtoi("-91283472332") << endl;
}
