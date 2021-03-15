#include "stdc++.h"
using namespace std;

// 215:第k大数

// Quick-Select
// class Solution {
// public:

//     int quickSelect(vector<int>& nums, int l, int r, int index) {
//         int q = randomPartition(nums, l, r);
//         if (q == index) {
//             return nums[q];
//         }
//         return (q < index) ? quickSelect(nums,q+1,r,index) : quickSelect(nums,l,q-1,index);
//     }

//     int randomPartition(vector<int>& nums, int l, int r) {
//         int i = rand() % (r-l+1) + l;
//         swap(nums[i], nums[r]);
//         return partition(nums,l,r);
//     }

//     int partition(vector<int>& nums, int l, int r){
//         int pivot = nums[r], i = l-1;
//         for (int j = l; j < r ; ++j){
//             if (nums[j] <= pivot) {
//                 i ++;
//                 swap(nums[i], nums[j]);
//             }
//         }
//         i ++;
//         swap(nums[i], nums[r]);
//         return i;
//     }

//     int findKthLargest(vector<int>& nums, int k) {
//         srand(time(0));
//         return quickSelect(nums,0,nums.size()-1, nums.size()-k);
//     }
// };

// 堆排序

class Solution {
public:
    void maxHeapify(vector<int>& nums,int i, int heapSize) {
        int l = 2*i+1, r = 2*i+2, largest = i;
        if (l < heapSize && nums[l] > nums[largest]) {
            largest = l;
        } 
        if (r < heapSize && nums[r] > nums[largest]) {
            largest = r;
        }
        if (largest != i) {
            swap(nums[i], nums[largest]);
            maxHeapify(nums, largest, heapSize);
        }
    }

    void buildMaxHeap(vector<int>& nums, int heapSize) {
        for (int i=heapSize/2;i>=0;--i) {
            maxHeapify(nums, i, heapSize);
        }
    }

    int findKthLargest(vector<int>& nums, int k) {
        int heapSize = nums.size();
        buildMaxHeap(nums, heapSize);
        for (int i = nums.size()-1; i >= nums.size()-k; i--) {
            swap(nums[0], nums[i]);
            --heapSize;
            maxHeapify(nums,0,heapSize);
        }
        return nums[0];
    }
};