package main

func twoSum(numbers []int, target int) []int {

	var res1, res2 int
	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] == target {
			res1, res2 = l+1, r+1
		} else {
			if numbers[l]+numbers[r] > target {
				r = r - 1
			} else {
				l = l + 1
			}
		}
	}

	return []int{res1, res2}
}
