#include <bits/stdc++.h> // for conveience
using namespace std;

/*
    INT_MIN需要特判
 */
class Solution {
public:
    int reverse(int x) {
        if (x == INT_MIN) return 0;
        bool fu = (x < 0)? true: false;
        if (fu) x = -x;
        long long res = 0;
        while (x > 0){
            res = res * 10 + (x % 10);
            x /= 10;
        }        
        if (fu) res = -res;
        if (res > INT_MAX || (fu && res < INT_MIN)){
            res = 0;
        }
        return res;
    }
};
