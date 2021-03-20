package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func restoreIPAddresses(s string) []string {
	var res []string
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rem := len(s) - (i + j + k)
				if rem < 1 || rem > 3 {
					continue
				}

				ip := []string{s[:i], s[i : i+j], s[i+j : i+j+k], s[i+j+k:]}

				// ip1, ip2, ip3, ip4 := s[:i], s[i:i+j], s[i+j:i+j+k], s[i+j+k:]

				ok := true
				for i := 0; i <= 3; i++ {
					ipval, _ := strconv.ParseInt(ip[i], 10, 16)
					fmt.Println(ipval)
					if ipval > 255 || notValidNumber(ip[i]) {
						ok = false
						break
					}
				}

				if !ok {
					continue
				}

				var buf bytes.Buffer
				for i := 0; i < 3; i++ {
					buf.WriteString(ip[i])
					buf.WriteByte('.')
				}
				buf.WriteString(ip[3])

				res = append(res, buf.String())
			}
		}
	}
	return res
}

func notValidNumber(s string) bool {
	return len(s) != 1 && s[0] == '0'
}
