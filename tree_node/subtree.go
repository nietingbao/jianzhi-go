package treenode

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here
	if pRoot2 == nil || pRoot1 == nil {
		return false
	}
	if pRoot1.Val == pRoot2.Val {
		var leftC, rightC bool
		if pRoot2.Left == nil {
			leftC = true
		} else if pRoot1.Left != nil {
			leftC = equalWithRoot(pRoot1.Left, pRoot2.Left)
		}
		if !leftC {
			return findSubInChild(pRoot1, pRoot2)
		}
		if pRoot2.Right == nil {
			rightC = true
		} else if pRoot1.Right != nil {
			rightC = equalWithRoot(pRoot1.Right, pRoot2.Right)
		}
		if rightC {
			return true
		} else {
			return findSubInChild(pRoot1, pRoot2)
		}
	} else {
		return findSubInChild(pRoot1, pRoot2)
	}
}

func equalWithRoot(par, chi *TreeNode) bool {
	if chi == nil {
		return true
	}
	if par == nil {
		return false
	}
	if par.Val != chi.Val {
		return false
	}
	var leftE, rightE bool
	if chi.Left == nil {
		leftE = true
	} else {
		if par.Left == nil {
			leftE = false
		} else {
			leftE = equalWithRoot(par.Left, chi.Left)
		}
	}
	if chi.Right == nil {
		rightE = true
	} else {
		if par.Right == nil {
			rightE = false
		} else {
			rightE = equalWithRoot(par.Right, chi.Right)
		}
	}
	return leftE && rightE
}

func findSubInChild(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	var res bool
	if pRoot1.Left != nil {
		res = HasSubtree(pRoot1.Left, pRoot2)
	}
	if !res {
		if pRoot1.Right != nil {
			return HasSubtree(pRoot1.Right, pRoot2)
		}
	} else {
		return true
	}
	return false
}

func PreOrder(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		if root.Left != nil {
			res = append(res, PreOrder(root.Left)...)
		}
		if root.Right != nil {
			res = append(res, PreOrder(root.Right)...)
		}
	}
	return res
}

func InOrder(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		if root.Left != nil {
			res = append(res, InOrder(root.Left)...)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			res = append(res, InOrder(root.Right)...)
		}
	}
	return res
}

func HasSubSlice(adult, child []int) bool {
	if len(child) == 0 {
		return false
	}
	for i, aNum := range adult {
		if aNum == child[0] {
			if len(child) == 1 {
				return true
			}
			ic := 1
			ia := i + 1
			for ic < len(child) && ia < len(adult) {
				if adult[ia] != child[ic] {
					break
				}
				ia++
				ic++
			}
			if ic == len(child) && ia <= len(adult) {
				return true
			}
		}
	}
	return false
}
