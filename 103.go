package main

/*
	102 & reverse
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	levelOrderRecur(root, 0, &res)
	for i := range res {
		if i%2 != 0 {
			leni := len(res[i])
			for j := 0; j < leni-j-1; j++ {
				res[i][j], res[i][leni-j-1] = res[i][leni-j-1], res[i][j]
			}
		}
	}
	return res
}
