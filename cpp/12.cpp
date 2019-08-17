#include <bits/stdc++.h> // for conveience
using namespace std;


/*
    一位位处理
 */

class Solution {
private:
    char one[4] = {'I','X','C','M'};
    char five[3] = {'V','L','D'};
public:
    string intToRoman(int num) {
        string res;
        stringstream ss;
        int ind = 0;
        while (num){
            int now = num % 10;
            if (now == 4){
                ss << five[ind];
                ss << one[ind]; 
            }
            else if (now == 9){
                ss << one[ind+1];
                ss << one[ind];
            }
            else {
                if (now >= 5){
                    for (int i=1;i<=max(0,now-5);++i){
                        ss << one[ind];
                    }
                    ss << five[ind];
                }
                else {
                    for (int i=1;i<=now;++i){
                        ss << one[ind];
                    }
                }
            }
            ++ind;
            num = num / 10;
        }
        ss >> res;
        reverse(res.begin(),res.end());
        return res;
    }
};