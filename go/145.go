package main

/*
	需要知道是从左子树退回到根节点，还是从右子树退回来的
*/

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var p *TreeNode
	stack := make([]*TreeNode, 100)
	flag, top := 0, -1
	if root != nil {
		top++
		stack[top] = root
		tmp := root.Left
		for top != -1 {
			for tmp != nil {
				top++
				stack[top] = tmp
				tmp = tmp.Left
			}
			p = nil
			flag = 1
			for top != -1 && flag == 1 {
				tmp = stack[top]
				if tmp.Right == p {
					res = append(res, tmp.Val)
					top--
					p = tmp
				} else {
					tmp = tmp.Right
					flag = 0
				}
			}
		}
	}

	return res
}
