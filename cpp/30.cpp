#include <bits/stdc++.h>
using namespace std;

/*
    单词的长度相同是关键
    通过在s中取一段totalLen的距离查看出现的单词次数来比较
 */
class Solution {
private:
    map<string,int> m;
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        vector<int> res;
        int wordSize = words.size();
        if (wordSize == 0) return res;
        int wordLen = words[0].length(),totalLen = wordSize * wordLen,sLen = s.length();
        for (string word : words){
            ++m[word];
        }
        map<string,int> mm;
        vector<int> omit(wordLen,-1);
        for (int i=0;i<wordLen;++i){
            int l = i, r = i, cnt = 0;
            mm.clear();
            while (r + wordLen <= sLen){
                string w = s.substr(r,wordLen);
                r += wordLen;
                if (m.count(w) == 0){
                    cnt = 0;
                    l = r;
                    mm.clear();
                }
                else {
                    ++mm[w];
                    ++cnt;
                    while (mm[w] > m[w]){
                        string lw = s.substr(l,wordLen);
                        --mm[lw];
                        --cnt;
                        l += wordLen;
                    }
                    if (cnt == wordSize) res.push_back(i);
                }
            }
        }
        return res;
    }
};

