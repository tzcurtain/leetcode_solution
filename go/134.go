package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	total := 0
	nowres := -1
	res := -1
	for i := 0; i < n; i++ {
		nowgas := gas[i] - cost[i]
		total += nowgas
		if res == -1 && nowgas >= 0 {
			res = i
			nowres = nowgas
		} else {
			nowres += nowgas
			if nowres < 0 {
				res = -1
				nowres = -1
			}
		}
	}

	if total <= 0 {
		return 0
	}

	return res + 1
}
