package main

import (
	"bytes"
)

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	now := 0
	wordLen := len(words)
	nowUsed, wordCount := 0, 0
	for now <= wordLen {
		if now == wordLen || nowUsed+len(words[now])+wordCount > maxWidth {
			//fmt.Println(now)
			totalComma := maxWidth - nowUsed
			avgComma, extraComma := totalComma, 0
			if wordCount != 1 {
				avgComma = totalComma / (wordCount - 1)
				extraComma = totalComma % (wordCount - 1)
			}
			if now == wordLen {
				avgComma = 1
				extraComma = 0
			}

			var writeBuf bytes.Buffer
			// fmt.Println("Comma", avgComma, extraComma)
			writeBuf.WriteString(words[now-wordCount])

			for i := now - wordCount + 1; i < now; i++ {
				if i-now+wordCount <= extraComma {
					writeBuf.WriteByte(' ')
				}
				for j := 1; j <= avgComma; j++ {
					writeBuf.WriteByte(' ')
				}
				writeBuf.WriteString(words[i])
			}
			if wordCount == 1 && now != wordLen {
				for i := 1; i <= avgComma; i++ {
					writeBuf.WriteByte(' ')
				}
			}

			if now == wordLen {
				remainLen := maxWidth - writeBuf.Len()
				for i := 1; i <= remainLen; i++ {
					writeBuf.WriteByte(' ')
				}
			}

			res = append(res, writeBuf.String())
			if now != wordLen {
				nowUsed = len(words[now])
				wordCount = 1
			}
		} else {
			nowUsed += len(words[now])
			wordCount++
		}
		now++
	}

	return res
}
