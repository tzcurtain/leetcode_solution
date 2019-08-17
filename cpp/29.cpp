#include <bits/stdc++.h>
using namespace std;

/*
    x / y = a ... b -> x = ay + b -> x = (2^0*a1+2^1*a2+...+2^n*an)*y + b
    位运算解决
 */

class Solution {
public:
    int divide(int dividend, int divisor) {
        vector<long long> divs;
        vector<int> ress;
        bool isNeg1 = dividend < 0;
        bool isNeg2 = divisor < 0;
        bool isPos = !((dividend > 0)^(divisor > 0));
        long long divsor = abs((long long)divisor);
        long long divend = abs((long long)dividend);
        long long now = divsor,res = 0;
        while (now <= divend){
            divs.push_back(now);
            ress.push_back(0);
            now = now << 1;
        }
        int size = divs.size();
        for (int i=size-1;i>=0;--i){
            while (divend >= divs[i]){
                divend -= divs[i]; 
                ++ress[i];
            }
        }
        long long base = 1;
        for (int i=0;i<size;++i){
            res += base * ress[i];
            base = base << 1;
        }
        if (isPos){
            if (res > INT_MAX) return INT_MAX;
            else return res;
        }
        else {
            if (res < INT_MIN) return INT_MAX;
            else return -res;
        }   
    }
};

int main(){
    Solution s = Solution();
    cout << (1^1) << endl;
    bool isNe = ((-222<0)^(-222<1));
    cout << isNe << endl;
    cout << s.divide(-2147483648,-1) << endl;
}