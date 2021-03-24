#include "stdc++.h"
using namespace std;

// 295: 数据流的中位数
// mulitiset


class MedianFinder {
private:
    multiset<int> m;
    multiset<int>::iterator lo, hi;
public:
    /** initialize your data structure here. */
    MedianFinder() {

    }
    
    void addNum(int num) {
        int n = m.size();
        m.insert(num);
        if (!n) {
            lo = hi = m.begin();
        } else if (n % 2 == 1){
            if (num < *lo) {
                lo --;
            } else {
                hi ++;
            }
        } else {
            if (num > *lo && num < *hi) {
                lo ++;
                hi --;
            } else if (num >= *hi) {
                lo ++;
            } else {
                hi --;
                lo = hi;
            }
        }
        n ++;
    }
    double findMedian() {
        return (*lo + *hi) * 0.5;
    }
};


// maxHeap & minHeap
// class MedianFinder {
// private:
//     priority_queue<int,vector<int>, greater<int> > hi;
//     priority_queue<int,vector<int>, less<int> > lo;
// public:
//     /** initialize your data structure here. */
//     MedianFinder() {
        
//     }
    
//     void addNum(int num) {
//         lo.push(num);
//         hi.push(lo.top());
//         lo.pop();
//         if (lo.size() < hi.size()) {
//             lo.push(hi.top());
//             hi.pop();
//         }
//     }
    
//     double findMedian() {
//         return (lo.size() > hi.size()) ? (double)lo.top() : (hi.top() + lo.top()) * 0.5;
//     }
// };


// hash表
// class MedianFinder {
// private:
//     map<int, int> m;
//     int size;
// public:
//     /** initialize your data structure here. */
//     MedianFinder() {
//         this->size = 0;
//     }
    
//     void addNum(int num) {
//         if (this->m.find(num) == this->m.end()) {
//             this->m[num] = 0;
//         } 
//         this->m[num] ++;
//         this->size ++;
//     }
    
//     double findMedian() {
//         double res;
//         bool flag = false;
//         int now = 0;
//         for (auto &it : m) {
//             now += it.second;
//             if (now < size / 2) { 
//                 continue;
//             }
//             if (now == size/2) {
//                 if (size % 2 == 0) {
//                     flag = true;
//                     res = it.first;
//                 }
//             }
//             if (now >= (size/2+1)) {
//                 if (size % 2 != 0) {
//                     return it.first;
//                 }
//                 if (flag) {
//                     return (res + it.first) * 0.5;
//                 } 
//                 return it.first;
//             }
//         }
//         return res;
//     }
// };