package main

func merge88(nums1 []int, m int, nums2 []int, n int) {
	i, j, now := 0, 0, 0
	nums3 := make([]int, m+n)
	for i < m && j < n {
		if nums1[i] > nums2[j] {
			nums3[now] = nums2[j]
			j++
		} else {
			nums3[now] = nums1[i]
			i++
		}
		now++
	}

	if i == m {
		for j < n {
			nums3[now] = nums2[j]
			j++
			now++
		}
	} else {
		for i < m {
			nums3[now] = nums1[i]
			i++
			now++
		}
	}

	copy(nums1, nums3)
}
