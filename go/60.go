package main

import (
	"bytes"
	"fmt"
	"sort"
)

/*
	根据规律直接找答案
	n * (n-1)!
*/

func getPermutation(n int, k int) string {
	tmp := make([]int, n)
	fac := make([]int, n+2)
	fac[0] = 1

	for i := range tmp {
		tmp[i] = i + 1
		fac[i+1] = fac[i] * (i + 1)
	}
	// fmt.Println(fac)
	var buf bytes.Buffer
	for i := 0; i < n-1; i++ {
		facnow := fac[n-i-1]
		swapNow := (k - 1) / facnow
		fmt.Println(swapNow, facnow)
		k -= swapNow * facnow
		tmp[i], tmp[i+swapNow] = tmp[i+swapNow], tmp[i]
		sort.Ints(tmp[i+1:])
		buf.WriteByte('0' + byte(tmp[i]))
		fmt.Println(tmp)
	}
	buf.WriteByte('0' + byte(tmp[n-1]))

	return buf.String()
}
