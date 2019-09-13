package main

import "bytes"

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	wordlen := len(words)
	bef := 0
	now := 1
	nowlen := len(words[0])
	for now < wordlen {
		if nowlen+len(words[now]) <= maxWidth {
			nowlen += len(words[now])
		} else {
			var buf bytes.Buffer

			res = append(res, buf.String())
			bef = now
			nowlen = len(words[now])
		}
		now++
	}
	var buf bytes.Buffer
	for i := bef; i < now; i++ {
		buf.WriteString(words[i])
	}
	res = append(res, buf.String())
	
	return res
}
