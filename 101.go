package main

func isSymmetric(root *TreeNode) bool {
	// return isSymmetricNonRecur(root)
	if root == nil {
		return true
	}
	return isSymmetricRecur(root.Left, root.Right)
}

func isSymmetricRecur(l, r *TreeNode) bool {
	if l == nil || r == nil {
		if l == nil && r == nil {
			return true
		}
		return false
	}

	return l.Val == r.Val && isSymmetricRecur(l.Left, r.Right) && isSymmetricRecur(l.Right, r.Left)

}

func isSymmetricNonRecur(root *TreeNode) bool {
	if root == nil {
		return true
	}
	st := make([]*TreeNode, 1000)
	size := 0
	if root.Left != nil {
		st[size] = root.Left
		size++
	}
	if root.Right != nil {
		st[size] = root.Right
		size++
	}

	for size != 0 {
		if size%2 != 0 {
			return false
		}
		nowl, nowr := st[size-1], st[size-2]
		// fmt.Println(nowl.Val, nowr.Val)
		size -= 2
		if nowl.Val != nowr.Val {
			return false
		}
		if nowl.Left == nil || nowr.Right == nil {
			if nowl.Left != nil || nowr.Right != nil {
				return false
			}
		} else {
			st[size] = nowl.Left
			size++
			st[size] = nowr.Right
			size++
		}

		if nowl.Right == nil || nowr.Left == nil {
			if nowl.Right != nil || nowr.Left != nil {
				return false
			}
		} else {
			st[size] = nowl.Right
			size++
			st[size] = nowr.Left
			size++
		}
	}

	return true
}
