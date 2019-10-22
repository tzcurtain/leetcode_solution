package main

/*
	reuse q102 code
*/

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	levelOrderRecur(root, 0, &res)

	lenres := len(res)
	for i := 0; i < lenres-i; i++ {
		res[i], res[lenres-i] = res[lenres-i], res[i]
	}

	return res
}
