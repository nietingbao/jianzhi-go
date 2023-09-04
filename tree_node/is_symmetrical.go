package treenode

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}
	if pRoot.Left != nil {
		if pRoot.Right != nil {
			return isMirror(pRoot.Left, pRoot.Right)
		} else {
			return false
		}
	} else {
		if pRoot.Right != nil {
			return false
		} else {
			return true
		}
	}
}

func isMirror(a, b *TreeNode) bool {
	if a.Val != b.Val {
		return false
	}
	var le, ri bool
	if a.Left != nil {
		if b.Right == nil {
			return false
		} else {
			le = isMirror(a.Left, b.Right)
		}
	} else {
		if b.Right == nil {
			le = true
		} else {
			return false
		}
	}
	if a.Right != nil {
		if b.Left == nil {
			return false
		} else {
			ri = isMirror(a.Right, b.Left)
		}
	} else {
		if b.Left == nil {
			ri = true
		} else {
			return false
		}
	}
	return le && ri
}
