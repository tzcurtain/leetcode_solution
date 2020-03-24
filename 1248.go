package main

func numberOfSubarrays(nums []int, k int) int {
	res := 0
	nowK := 0
	lastEven, lastOdd := 0, -1
	for i := range nums {
		if nums[i]%2 != 0 {
			lastOdd = i
			break
		}
	}

	if lastOdd == -1 {
		return 0
	}

	for i := range nums {
		if nums[i]%2 == 0 {
			if nowK == k {
				res += lastOdd - lastEven + 1
			}
		} else {
			nowK++
			if nowK == k {
				res += lastOdd - lastEven + 1
			} else if nowK == k+1 {
				tmp := lastOdd
				lastOdd++
				for lastOdd < len(nums) && nums[lastOdd]%2 == 0 {
					lastOdd++
				}
				lastEven = tmp + 1
				res += lastOdd - lastEven + 1
				nowK = k
			}
		}
	}
	return res
}
