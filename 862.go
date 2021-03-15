package main

/* waiting for finish */

func shortestSubarray(A []int, K int) int {
	if len(A) == 0 {
		return -1
	}
	res := -1
	l, r, now, n := 0, 0, A[0], len(A)

	for r < n {
		if now >= K {
			if res == -1 {
				res = r - l + 1
			} else {
				res = min(res, r-l+1)
			}
			if l < r {
				l++
				now -= A[l]
			} else { // l == r
				return 1
			}
		} else {
			r++
			if r == n {
				break
			}
			now += A[r]
		}
	}

	return -1
}
