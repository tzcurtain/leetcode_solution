package main

import (
	"sort"
)

/*
	排序后用map筛选
*/
type byteSlice []byte

func (a byteSlice) Len() int {
	return len(a)
}

func (a byteSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byteSlice) Less(i, j int) bool {
	return a[i] < a[j]
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, val := range strs {
		sArr := byteSlice(val)
		sort.Sort(sArr)
		sNow := string(sArr[:])
		// fmt.Println(sNow)
		x := m[sNow]
		x = append(x, val)
		m[sNow] = x
	}

	var res [][]string
	for _, sArrNow := range m {
		res = append(res, sArrNow)
	}

	return res
}
