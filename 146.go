package main

type LRUCache struct {
	m        map[int]*DoubleListNode
	cap      int
	size     int
	listEnd  *DoubleListNode
	listHead *DoubleListNode
}

type DoubleListNode struct {
	Next *DoubleListNode
	Bef  *DoubleListNode
	Val  int
	Key  int
}

func Constructor(capacity int) LRUCache {
	var l LRUCache
	l.cap = capacity
	l.m = make(map[int]*DoubleListNode, capacity)
	l.size = 0
	l.listEnd = nil
	l.listHead = nil
	return l
}

func (this *LRUCache) Get(key int) int {
	val := this.m[key]
	if val != nil {
		// put val in the end
		if this.listHead == this.listEnd || this.listEnd == val { // one only or it's listEnd already
			return val.Val
		}
		befNode, nexNode := val.Bef, val.Next
		if this.listHead == val {
			this.listHead = nexNode
		}
		if befNode != nil {
			befNode.Next = nexNode
		}
		if nexNode != nil {
			nexNode.Bef = befNode
		}
		val.Bef, val.Next = this.listEnd, nil
		this.listEnd.Next = val
		this.listEnd = val
		return val.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.m[key].Val = value
		return
	}
	if this.size == this.cap {
		// do sth
		keyTodelete := this.listHead.Key
		if this.listEnd == this.listHead { // caps = 1
			this.listEnd = nil
		}
		this.listHead = this.listHead.Next
		delete(this.m, keyTodelete)
		this.size--
	}
	this.size++
	var tmp *DoubleListNode = new(DoubleListNode)
	tmp.Val, tmp.Key = value, key
	tmp.Next, tmp.Bef = nil, this.listEnd
	if this.listHead == nil {
		this.listHead = tmp
	}
	if this.listEnd != nil {
		this.listEnd.Next = tmp
	}
	this.listEnd = tmp
	this.m[key] = tmp
}
