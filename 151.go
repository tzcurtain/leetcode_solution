package main

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	l, r := -1, -1

	var buf bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' || i == 0 {
			if r != -1 {
				l = i + 1
				if i == 0 {
					l = 0
				}
				buf.WriteString(s[l : r+1])
				buf.WriteByte(' ')
				r = -1
			}
			continue
		}
		if r == -1 {
			r = i
		}
	}
	if s[0] != ' ' && s[1] == ' ' {
		buf.WriteByte(s[0])
	}
	res := buf.String()
	if len(res) > 0 && res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}
	return res
}
