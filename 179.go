package main

import (
	"bytes"
	"sort"
	"strconv"
)

type myString struct {
	a string
}

type myStringSlice []myString

func (p myStringSlice) Less(i, j int) bool {
	return (p[i].a + p[j].a) > (p[j].a + p[i].a)
}

func (p myStringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p myStringSlice) Len() int {
	return len(p)
}

func largestNumber(nums []int) string {
	var strNums myStringSlice
	for i := range nums {
		strNums = append(strNums, myString{strconv.Itoa(nums[i])})
	}

	sort.Sort(strNums)

	var buf bytes.Buffer
	for i := range strNums {
		buf.WriteString(strNums[i].a)
	}

	res := buf.String()

	now := 0
	for now < len(res)-1 && res[now] == '0' {
		now++
	}
	res = res[now:]

	return res
}
