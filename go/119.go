package main

func getRow(numRows int) []int {
	res := make([]int, numRows+1)
	res[0] = 1
	for i := 0; i <= numRows; i++ {
		for j := i; j >= 1; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
