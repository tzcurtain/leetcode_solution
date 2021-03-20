/*
 * @lc app=leetcode.cn id=1504 lang=golang
 *
 * [1504] 统计全 1 子矩形
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSubmat(mat [][]int) int {
	row, col := len(mat), len(mat[0])
	nums := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			for ilen := 1; ilen <= row; ilen++ {
				rdi := i + ilen - 1
				if rdi >= row {
					break
				}
				for jlen := 1; jlen <= col; jlen++ {
					rdj := j + jlen - 1
					if rdj >= col {
						break
					}
					ok := true
					for ii := i; ii <= rdi; ii++ {
						if mat[ii][rdj] == 0 {
							ok = false
							break
						}
					}
					if !ok {
						break
					}
					nums++
				}
			}
		}
	}
	return nums
}

// @lc code=end

