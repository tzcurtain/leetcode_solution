package main

// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
// type BSTIterator struct {
// 	nowVal   []int
// 	nowIndex int
// }

// /** construct */
// func Constructor(root *TreeNode) BSTIterator {
// 	var nowVal []int
// 	var nowSt []*TreeNode
// 	top := -1
// 	for root.Left != nil {
// 		root = root.Left
// 		nowSt = append(nowSt, root)
// 		top++
// 	}

// 	for top != -1 {
// 		now := nowSt[top]
// 	}

// 	return BSTIterator{nowVal, -1}
// }

// func (this *BSTIterator) Next() int {
// 	this.nowIndex++
// 	return this.nowVal[this.nowIndex]
// }

// func (this *BSTIterator) HasNext() bool {
// 	return this.nowIndex < len(this.nowVal)
// }

// /**
//  * Your BSTIterator object will be instantiated and called as such:
//  * obj := Constructor(root);
//  * param_1 := obj.Next();
//  * param_2 := obj.HasNext();
//  */
