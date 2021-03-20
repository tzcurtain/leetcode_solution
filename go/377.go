package main

/*
	f[i] = f[i-nums[j]] + nums[j] (i = 1..target, j = 1...i)
*/

func sliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func combinationSum4(nums []int, target int) int {
	f := make([]int, target+2)
	for i := 1; i <= target; i++ {
		for _, j := range nums {
			if i == j {
				f[i]++
			} else if i > j {
				f[i] += f[i-j]
			}
			// fmt.Println(i, ":", m[i])
		}
	}

	return f[target]
}

// func main() {
// 	candidates := []int{1, 50}
// 	target := 200
// 	fmt.Println(combinationSum4(candidates, target))
// }
