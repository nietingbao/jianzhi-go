package treenode

func TreeDepth(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}
	tr, tl := TreeDepth(pRoot.Right), TreeDepth(pRoot.Left)
	if tr >= tl {
		return tr + 1
	} else {
		return tl + 1
	}
}
