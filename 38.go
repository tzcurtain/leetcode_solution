package main

import (
	"bytes"
)

func countAndSay(n int) string {
	res := "1"
	for i := 2; i <= n; i++ {
		var buffer bytes.Buffer
		now := 1
		for idx := range res {
			//fmt.Println(i, " ", idx, " ", now, " ", buffer.String())
			if idx+1 < len(res) && res[idx] == res[idx+1] {
				now++
				continue
			}
			buffer.WriteByte('0' + byte(now))
			buffer.WriteByte(res[idx])
			//fmt.Println(i, " ", idx, " ", buffer.String())
			now = 1
		}

		res = buffer.String()
	}

	return res
}

// func main() {
// 	fmt.Println(countAndSay(30))
// }
