package main

import "math"

/*
	const
*/
const (
	Modint = 1000000007
)

/*
ListNode type definition
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/* 迷宫方向常量 */
const (
	SpiralRight = iota
	SpiralDown
	SpiralLeft
	SpiralUp
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPrime(a int) bool {
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func isAlpha(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}
