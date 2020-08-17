package main

import "bytes"

// Reverse string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func convertToTitle(n int) string {
	var buf bytes.Buffer
	for n != 0 {
		now := n % 26
		n = n / 26
		buf.WriteByte(byte(now) + byte(1) + 'A')
	}
	return Reverse(buf.String())
}
