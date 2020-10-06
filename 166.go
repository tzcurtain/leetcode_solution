package main

import (
	"bytes"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {

	fu := false

	if numerator < 0 {
		if denominator > 0 {
			fu = true
		}
	} else if denominator < 0 {
		fu = true
	}

	var buffer bytes.Buffer

	var num, denum int64
	num, denum = int64(numerator), int64(denominator)
	if num < 0 {
		num = -num
	}
	if denum < 0 {
		denum = -denum
	}

	if fu {
		buffer.WriteByte('-')
	}

	if num > denum {
		div := num / denum
		num -= div * denum
		buffer.WriteString(strconv.Itoa(int(div)))
	}

	m := make(map[int]int, 30)

	now := 1
	m[int(num)] = now

	// 整数部分
	if buffer.Len() == 0 {
		buffer.WriteByte('0')
	}

	if num != 0 {
		buffer.WriteByte('.')
	}

	for num != 0 {
		for num*10 < denum {
			num *= 10
			buffer.WriteByte('0')
			now++
		}
		num *= 10
		now++
		div := num / denum
		num -= div * denum
		// fmt.Println(numerar)
		buffer.WriteString(strconv.Itoa(int(div)))
		if m[int(num)] == 0 {
			m[int(num)] = now
		} else { // find
			// 小数点后(m[numerator]) - (now-1)位循环
			res := buffer.String()

			// fmt.Println(m[numerator], numerator, now, len(res))
			l := m[numerator] - now + len(res)

			// recycle := res[l]

			var buffer2 bytes.Buffer

			buffer2.WriteString(res[:l])
			buffer2.WriteByte('(')
			buffer2.WriteString(res[l:])
			buffer2.WriteByte(')')
			//fmt.Println("REcycle", recycle)
			return buffer2.String()
		}

	}

	return buffer.String()
}
