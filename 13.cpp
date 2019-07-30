#include <bits/stdc++.h> // for conveience
using namespace std;


/*
    一位位处理
 */

class Solution {
private:
    int val[200];
public:
    int romanToInt(string s) {
        val['I'] = 1; val['V'] = 5; val['X'] = 10; val['L'] = 50;
        val['C'] = 100; val['D'] = 500; val['M'] = 1000;
        int sum = val[s[0]];
        for(int i=1; i<s.length(); i++){
            int temp = 0;
            //如果当前指向的字符比前一位字符对应的数字要大，则要对二者作差
            if(val[s[i]] > val[s[i-1]]){
                temp = val[s[i]] - 2 * val[s[i-1]];
            }
            //如果当前指向的字符比前一位字符所对应的数字要小，则将当前字符对应的数加入sum
            else if(val[s[i]] <= val[s[i-1]]){
                temp = val[s[i]];
            }
            sum += temp;
        }
        return sum;    
    }
};