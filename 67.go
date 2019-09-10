package main

/*
	高精度加法二进制版
*/
func addBinary(a string, b string) string {
	lena, lenb := len(a), len(b)
	res := make([]byte, max(lena, lenb))
	i, j, k := lena-1, lenb-1, len(res)-1
	for i >= 0 && j >= 0 {
		res[k] += (a[i] - '0') + b[j]
		if res[k] >= '2' && k > 0 {
			res[k-1]++
			res[k] -= 2
		}
		k--
		i--
		j--
	}
	for i >= 0 {
		res[k] += a[i]
		if res[k] >= '2' && k > 0 {
			res[k-1]++
			res[k] -= 2
		}
		k--
		i--
	}
	for j >= 0 {
		res[k] += b[j]
		if res[k] >= '2' && k > 0 {
			res[k-1]++
			res[k] -= 2
		}
		k--
		j--
	}

	if res[0] >= '2' {
		res[0] -= 2
		res = append(res[:0], append([]byte{'1'}, res[0:]...)...)
	}

	return string(res)
}
