package main

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}
	l, r := 0, k
	total := 0
	i := 0
	for total < n {
		for i != r && total < n {
			nums[i], nums[r] = nums[r], nums[i]
			total++
			l = r
			r = (l + k) % n
			if i == r {
				total++
			}
		}
		i++
		l, r = i, (i+k)%n
	}
}
