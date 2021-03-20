package main

// lp 左边节点通过当前节点的最大路径， rp 右边节点通过当前节点的最大路径
// l , r 左子节点的res和右子节点的res

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, _, _, _, _ := maxPathSumRecur(root)
	return res
}

func maxPathSumRecur(root *TreeNode) (res, lp, rp, l, r int) {
	if root.Left == nil && root.Right == nil {
		res, lp, rp, l, r = root.Val, root.Val, root.Val, IntMin, IntMin
		return
	}

	if root.Left == nil {
		_, llp, rrp, ll, rr := maxPathSumRecur(root.Right)
		lp, rp, l, r = root.Val, max(root.Val, max(llp, rrp)+root.Val), IntMin, max(ll, max(rr, llp+rrp-root.Right.Val))
		res = max(lp+rp-root.Val, max(l, r))
		return
	} else if root.Right == nil {
		_, llp, rrp, ll, rr := maxPathSumRecur(root.Left)
		lp, rp, l, r = max(root.Val, max(llp, rrp)+root.Val), root.Val, max(ll, max(rr, llp+rrp-root.Left.Val)), IntMin
		res = max(lp+rp-root.Val, max(l, r))
		return
	}

	_, llp, rrp, ll, rr := maxPathSumRecur(root.Left)
	_, lllp, rrrp, lll, rrr := maxPathSumRecur(root.Right)

	lp, rp, l, r = max(root.Val, max(llp, rrp)+root.Val), max(root.Val, max(lllp, rrrp)+root.Val),
		max(ll, max(rr, llp+rrp-root.Left.Val)), max(lll, max(rrr, lllp+rrrp-root.Right.Val))

	res = max(lp+rp-root.Val, max(l, r))

	return
}
