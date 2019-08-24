package main

import (
	"bytes"
	"fmt"
)

/*
	高精度乘法 注意进位和最高位的处理
*/
func multiply(num1 string, num2 string) string {
	num3 := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		x1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			x2 := num2[j] - '0'
			pos := (len(num2) - 1 - j) + (len(num1) - 1 - i)
			res := x1 * x2
			fmt.Println(res, pos)
			num3[pos] += res
			num3[pos+1] += num3[pos] / 10
			num3[pos] = num3[pos] % 10
		}
	}

	now := len(num1) + len(num2) - 1
	for num3[now] == 0 && now > 0 {
		now--
	}
	for num3[now] > 10 {
		num3[now+1] = num3[now] / 10
		num3[now] = num3[now] % 10
		now++
	}

	var buffer bytes.Buffer

	for ; now >= 0; now-- {
		buffer.WriteByte(num3[now] + '0')
	}

	return buffer.String()
}
