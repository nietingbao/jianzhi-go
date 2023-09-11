package treenode

func lowestCommonAncestorS(root *TreeNode, p int, q int) int {
	// write code here
	if root == nil {
		return 0
	}
	r := root.Val
	if r == p {
		return p
	}
	if r == q {
		return q
	}
	if inDifferentSide(r, p, q) {
		return r
	}
	if bothLeft(r, p, q) {
		return lowestCommonAncestorS(root.Left, p, q)
	}
	if bothRight(r, p, q) {
		return lowestCommonAncestorS(root.Right, p, q)
	}
	return 0
}

func inDifferentSide(r, p, q int) bool {
	return p < r && q > r || p > r && q < r
}

func bothLeft(r, p, q int) bool {
	return p < r && q < r
}

func bothRight(r, p, q int) bool {
	return p > r && q > r
}
