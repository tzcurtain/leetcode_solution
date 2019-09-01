package main

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	res := 0
	nowSum := 0
	for i := range calories {
		if i <= k-1 {
			nowSum += calories[i]
			if i == k-1 {
				if nowSum < lower {
					res--
				} else if nowSum > upper {
					res++
				}
			}
			continue
		}
		nowSum += calories[i] - calories[i-k]
		if nowSum < lower {
			res--
		} else if nowSum > upper {
			res++
		}
	}

	return res
}
