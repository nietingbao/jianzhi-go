package treenode

func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return nil
	}
	right := pRoot.Right
	if pRoot.Left != nil {
		pRoot.Right = Mirror(pRoot.Left)
	} else {
		pRoot.Right = nil
	}
	if right != nil {
		pRoot.Left = Mirror(right)
	} else {
		pRoot.Left = nil
	}
	return pRoot
}
