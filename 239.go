package main

// import "fmt"

// // Deque 双端队列
// type Deque struct {
// 	head, tail *DListNode
// 	totalSize  int
// }

// func (dq *Deque) empty() bool {
// 	return dq.totalSize == 0
// }

// func (dq *Deque) pushFront(n int) {
// 	currentNode := dq.head
// 	newNode := new(DListNode)
// 	newNode.Val = n
// 	if (currentNode == nil)
// 		dq.tail = &newNode
// 		dq.head = &newNode
// 	} else {
// 		newNode.Next = currentNode
// 		currentNode.Next = newNode
// 		dq.head = newNode
// 	}
// 	dq.head = dq.head.Next
// }
// func (dq *Deque) pushBack(n int) {
// 	if dq.empty() {
// 		dq.head = new(DListNode)
// 		dq.head.Val = n
// 		dq.tail = dq.head
// 		fmt.Println("WHY", dq.head == nil, dq.tail == nil, dq.empty())
// 		return
// 	}
// 	fmt.Println("EHLLLLL")
// 	dq.tail.Bef = new(DListNode)
// 	dq.tail.Bef.Val = n
// 	dq.tail.Bef.Next = dq.tail
// 	dq.tail = dq.tail.Bef
// }
// func (dq *Deque) popFront() {
// 	if dq.head.Bef != nil {
// 		dq.head = dq.head.Bef
// 		dq.head.Next = nil
// 	} else {
// 		dq.head = nil
// 		dq.tail = nil
// 	}
// }
// func (dq *Deque) popBack() {
// 	if dq.tail.Next != nil {
// 		dq.tail = dq.tail.Next
// 		dq.tail.Bef = nil
// 	} else {
// 		dq.head = nil
// 		dq.tail = nil
// 	}
// }
// func (dq *Deque) front() int {
// 	return dq.head.Val
// }
// func (dq *Deque) back() int {
// 	return dq.tail.Val
// }

// // MonoTonicQueue 单调队列
// type MonoTonicQueue struct {
// 	*Deque
// }

// func (mq MonoTonicQueue) push(n int) {
// 	for !mq.empty() && mq.back() < n {
// 		mq.popBack()
// 	}
// 	mq.pushBack(n)
// }
// func (mq MonoTonicQueue) max() int {
// 	return mq.front()
// }
// func (mq MonoTonicQueue) pop(n int) {
// 	if !mq.empty() && mq.front() == n {
// 		mq.popFront()
// 	}
// }

// func maxSlidingWindow(nums []int, k int) []int {
// 	mq := new(MonoTonicQueue)
// 	var res []int
// 	for i := 0; i < len(nums); i++ {
// 		if i < k-1 {
// 			mq.push(nums[i])
// 		} else {
// 			mq.push(nums[i])
// 			fmt.Println(mq.front())
// 			res = append(res, mq.max())
// 			mq.pop(nums[i-k+1])
// 		}
// 	}
// 	return res
// }
