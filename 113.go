package main

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var tmp []int
	nowsum := 0
	pathSumRecur(root, sum, nowsum, &res, tmp)
	return res
}

func pathSumRecur(root *TreeNode, sum, nowsum int, res *[][]int, tmp []int) {
	if root == nil {
		return
	}

	nowsum += root.Val
	tmp = append(tmp, root.Val)

	if nowsum == sum {
		if root.Left == nil && root.Right == nil {
			ttmp := make([]int, len(tmp))
			copy(ttmp, tmp)
			*res = append(*res, ttmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	pathSumRecur(root.Left, sum, nowsum, res, tmp)
	pathSumRecur(root.Right, sum, nowsum, res, tmp)

	if len(tmp) >= 1 {
		tmp = tmp[:len(tmp)-1]
	}

}
