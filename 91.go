package main

/*
	递推，注意0的情况
*/

func numDecodings(s string) int {
	lens := len(s)
	f := make([]int, lens)
	if s[0] != '0' {
		f[0] = 1
	}
	if lens == 1 || s[0] == '0' {
		return f[0]
	}
	if s[1] != '0' {
		f[1] = 1
	} else if s[1] > '2' {
		return 0
	}
	if s[0] == '1' || (s[0] == '2' && s[1] <= '6') {
		f[1]++
	}
	for i := 2; i < lens; i++ {
		if s[i] != '0' {
			f[i] = f[i-1]
		} else {
			if s[i-1] == '1' || s[i-1] == '2' {
				f[i] = f[i-2]
				continue
			} else {
				return 0
			}
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			f[i] += f[i-2]
		}
	}
	//fmt.Println(f)
	return f[lens-1]
}
