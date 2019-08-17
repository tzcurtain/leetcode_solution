#include <bits/stdc++.h> // for conveience
using namespace std;

/*
    通过将找中位数变为找第(n+m)/2大的数，然后通过二分不断缩小范围
 */
class Solution {
public:
    double findnth(vector<int>& nums1,vector<int>& nums2,int l1,int r1,int l2,int r2,int n){
        int len1 = r1-l1+1,len2 = r2-l2+1;
        if (len1 > len2){
            return findnth(nums2,nums1,l2,r2,l1,r1,n);
        }
        if (len1 == 0){
            return nums2[l2 + n - 1];
        }
        if (n == 1){
            return min(nums1[l1],nums2[l2]);
        }
        int i1 = l1 + min(len1,n/2)-1,i2 = l2+min(len2,n/2)-1;
        if (nums1[i1] > nums2[i2]){
            return findnth(nums1,nums2,l1,r1,i2+1,r2,n-(i2-l2+1));
        } 
        else {
            return findnth(nums1,nums2,i1+1,r1,l2,r2,n-(i1-l1+1));
        }
    }
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int n=nums1.size(),m=nums2.size();
        if ((n+m) % 2 != 0)
            return findnth(nums1,nums2,0,n-1,0,m-1,(n+m)/2+1);
        else
            return (findnth(nums1,nums2,0,n-1,0,m-1,(n+m)/2)
            +findnth(nums1,nums2,0,n-1,0,m-1,(n+m)/2+1))/2.0;
    }

};

int main(){
    Solution a;
    vector<int> a1 = {1,2},a2 = {3,4};
    // cout << "HH";
    cout << a.findMedianSortedArrays(a1,a2);
}